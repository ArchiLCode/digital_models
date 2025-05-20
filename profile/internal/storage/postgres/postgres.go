package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"profile/internal/domain/models"
	"profile/internal/storage"
	"strings"
	"time"

	"github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

// App возвращает приложение по id из базы данных PostgreSQL.
func (s *Storage) App(ctx context.Context, id int) (models.App, error) {
	const op = "storage.postgres.App"

	query := `
		SELECT id, name, secret
		FROM apps
		WHERE id = $1
	`

	// Выполняем запрос и сканируем результат в структуру models.App
	var app models.App
	err := s.db.QueryRowContext(ctx, query, id).Scan(&app.ID, &app.Name, &app.Secret)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.App{}, fmt.Errorf("%s: %w", op, storage.ErrAppNotFound)
		}

		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	return app, nil
}

func (s *Storage) CreateProfile(ctx context.Context, name, lastName, userID string) error {
	const op = "storage.postgres.CreateProfile"

	query := `
	INSERT INTO profiles (id, name, last_name)
	VALUES ($1, $2, $3)
	`

	_, err := s.db.ExecContext(ctx, query, userID, name, lastName)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key`) { // Код ошибки для нарушения уникальности
			return fmt.Errorf("%s: %w", op, storage.ErrUserExists)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) DeleteProfile(ctx context.Context, id string) error {
	const op = "storage.postgres.DeleteProfile"

	// Начинаем транзакцию
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: начало транзакции: %w", op, err)
	}

	// Функция для отката транзакции в случае ошибки
	defer func() {
		if err != nil {
			// Если произошла ошибка, откатываем транзакцию
			tx.Rollback()
		}
	}()

	queryPhoto := `
    DELETE FROM photos WHERE profile_id = $1
    `
	_, err = tx.ExecContext(ctx, queryPhoto, id)
	if err != nil {
		return fmt.Errorf("%s: удаление фотографий: %w", op, err)
	}

	// Удаляем профиль
	query := `
    DELETE FROM profiles
    WHERE id = $1
    `
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s: удаление профиля: %w", op, err)
	}

	// Фиксируем транзакцию
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%s: фиксация транзакции: %w", op, err)
	}

	return nil
}

func (s *Storage) GetProfileByID(ctx context.Context, id string) (models.GetProfileByIDResponse, error) {
	const op = "storage.postgres.GetProfileByID"

	// Начинаем транзакцию
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return models.GetProfileByIDResponse{}, fmt.Errorf("%s: начало транзакции: %w", op, err)
	}

	// Функция для отката транзакции в случае ошибки
	defer func() {
		if err != nil {
			// Если произошла ошибка, откатываем транзакцию
			tx.Rollback()
		}
	}()

	query := `
		SELECT name, last_name, sex, date_of_birth, city, height, chest_size, waist_size, hip_size, weight, avatar_url
		FROM profiles
		WHERE id = $1
	`

	var resp models.GetProfileByIDResponse
	var (
		dob       sql.NullTime
		sex       sql.NullString
		city      sql.NullString
		height    sql.NullInt32
		chestSize sql.NullInt32
		waistSize sql.NullInt32
		hipSize   sql.NullInt32
		weight    sql.NullInt32
		avatarURL sql.NullString
		name      sql.NullString
		lastName  sql.NullString
	)

	err = tx.QueryRowContext(ctx, query, id).Scan(
		&name,
		&lastName,
		&sex,
		&dob,
		&city,
		&height,
		&chestSize,
		&waistSize,
		&hipSize,
		&weight,
		&avatarURL,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.GetProfileByIDResponse{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
		return models.GetProfileByIDResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	// Присваиваем значения с учетом NULL
	if name.Valid {
		resp.Name = name.String
	}
	if lastName.Valid {
		resp.LastName = lastName.String
	}
	if sex.Valid {
		resp.Sex = sex.String
	}
	if dob.Valid {
		resp.DateOfBirth = timestamppb.New(dob.Time)
	}
	if city.Valid {
		resp.City = city.String
	}
	if height.Valid {
		resp.Height = height.Int32
	}
	if chestSize.Valid {
		resp.ChestSize = chestSize.Int32
	}
	if waistSize.Valid {
		resp.WaistSize = waistSize.Int32
	}
	if hipSize.Valid {
		resp.HipSize = hipSize.Int32
	}
	if weight.Valid {
		resp.Weight = weight.Int32
	}
	if avatarURL.Valid {
		resp.AvatarURL = avatarURL.String
	}

	// Получаем фотографии
	queryPhoto := `
    SELECT photo_url FROM photos WHERE profile_id = $1 ORDER BY position
    `
	rows, err := tx.QueryContext(ctx, queryPhoto, id)
	if err != nil {
		return models.GetProfileByIDResponse{}, fmt.Errorf("%s: запрос фотографий: %w", op, err)
	}
	defer rows.Close()

	var photos []string
	for rows.Next() {
		var photo string
		err = rows.Scan(&photo)
		if err != nil {
			return models.GetProfileByIDResponse{}, fmt.Errorf("%s: чтение фотографии: %w", op, err)
		}
		photos = append(photos, photo)
	}

	// Проверка ошибок при итерации
	if err = rows.Err(); err != nil {
		return models.GetProfileByIDResponse{}, fmt.Errorf("%s: итерация по фотографиям: %w", op, err)
	}

	resp.Photos = photos
	// Фиксируем транзакцию
	if err = tx.Commit(); err != nil {
		return models.GetProfileByIDResponse{}, fmt.Errorf("%s: фиксация транзакции: %w", op, err)
	}
	return resp, nil
}

func (s *Storage) UpdateProfile(ctx context.Context, req models.UpdateProfileRequest) error {
	// Начинаем транзакцию
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()
	// Поиск существующего профиля
	var (
		profileID       string
		avatarURLNull sql.NullString
	)
	err = tx.QueryRowContext(ctx, `
        SELECT id, avatar_url FROM profiles WHERE id = $1
    `, req.UserID).Scan(&profileID, &avatarURLNull)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Если профиль не найден, создаем новый
			var dateOfBirth sql.NullTime
			if !req.DateOfBirth.IsZero() {
				dateOfBirth = sql.NullTime{Time: req.DateOfBirth, Valid: true}
			}

			// Используем универсальную функцию для определения avatar_url
			avatarURL := getAvatarURL(req.AvatarURL, req.Photos)

			err = tx.QueryRowContext(ctx, `
                INSERT INTO profiles (
                    id, name, last_name, sex, date_of_birth, city, height, chest_size, waist_size, hip_size, weight, avatar_url
                ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
                RETURNING id
            `, req.UserID, req.Name, req.LastName, req.Sex, dateOfBirth, req.City, req.Height, req.ChestSize, req.WaistSize, req.HipSize, req.Weight, avatarURL).Scan(&profileID)
			if err != nil {
				return fmt.Errorf("failed to create profile: %w", err)
			}
		} else {
			return fmt.Errorf("failed to fetch profile: %w", err)
		}
	} else {
		// Обновляем существующий профиль
		var dateOfBirth sql.NullTime
		if !req.DateOfBirth.IsZero() {
			dateOfBirth = sql.NullTime{Time: req.DateOfBirth, Valid: true}
		}
		// Используем универсальную функцию для определения avatar_url
		currentAvatarURL := ""
		if avatarURLNull.Valid {
			currentAvatarURL = avatarURLNull.String
		}
		avatarURL := getAvatarURL(currentAvatarURL, req.Photos)

		_, err = tx.ExecContext(ctx, `
            UPDATE profiles SET
                name = $1,
                last_name = $2,
                sex = $3,
                date_of_birth = $4,
                city = $5,
                height = $6,
                chest_size = $7,
                waist_size = $8,
                hip_size = $9,
                weight = $10,
                avatar_url = $11
            WHERE id = $12
        `, req.Name, req.LastName, req.Sex, dateOfBirth, req.City, req.Height, req.ChestSize, req.WaistSize, req.HipSize, req.Weight, avatarURL, profileID)
		if err != nil {
			return fmt.Errorf("failed to update profile: %w", err)
		}
	}
	// Удаляем старые фотографии
	if len(req.Photos) > 0 {
		// Если есть новые фотографии, удаляем те, которых нет в новом списке
		_, err = tx.ExecContext(ctx, `
            DELETE FROM photos WHERE profile_id = $1 AND photo_url != ALL($2)
        `, profileID, pq.Array(req.Photos))
		if err != nil {
			return fmt.Errorf("failed to delete old photos: %w", err)
		}
	} else {
		// Если список фотографий пуст, удаляем все фотографии профиля
		_, err = tx.ExecContext(ctx, `
            DELETE FROM photos WHERE profile_id = $1
        `, profileID)
		if err != nil {
			return fmt.Errorf("failed to delete all photos: %w", err)
		}
	}

	// Добавляем или обновляем фотографии
	for i, url := range req.Photos {
		var photoID uint
		err = tx.QueryRowContext(ctx, `
            SELECT id FROM photos WHERE profile_id = $1 AND photo_url = $2
        `, profileID, url).Scan(&photoID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				// Если фотография не найдена, создаем новую
				_, err = tx.ExecContext(ctx, `
                    INSERT INTO photos (profile_id, photo_url, position)
                    VALUES ($1, $2, $3)
                `, profileID, url, i)
				if err != nil {
					return fmt.Errorf("failed to insert photo: %w", err)
				}
			} else {
				return fmt.Errorf("failed to fetch photo: %w", err)
			}
		} else {
			// Если фотография найдена, обновляем её позицию
			_, err = tx.ExecContext(ctx, `
                UPDATE photos SET position = $1 WHERE id = $2
            `, i, photoID)
			if err != nil {
				return fmt.Errorf("failed to update photo: %w", err)
			}
		}
	}

	// Завершаем транзакцию
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *Storage) GetCatalog(ctx context.Context, req models.GetCatalogRequest) (models.GetCatalogResponse, error) {
	// Проверяем обязательные поля пагинации
	if req.Limit <= 0 {
		return models.GetCatalogResponse{}, fmt.Errorf("limit must be positive")
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	offset := (req.Page - 1) * req.Limit

	// Динамически строим WHERE условия
	var whereClauses []string
	var filterArgs []interface{}
	argPos := 1

	// Добавляем условия только для ненулевых фильтров
	var zeroTime time.Time
	if req.DateFrom != zeroTime {
		whereClauses = append(whereClauses, fmt.Sprintf("date_of_birth <= $%d", argPos))
		filterArgs = append(filterArgs, req.DateFrom)
		argPos++
	}
	if req.DateTo != zeroTime {
		whereClauses = append(whereClauses, fmt.Sprintf("date_of_birth > $%d", argPos))
		filterArgs = append(filterArgs, req.DateTo)
		argPos++
	}
	if req.Sex != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("sex = $%d", argPos))
		filterArgs = append(filterArgs, req.Sex)
		argPos++
	}
	if req.HeightFrom > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("height >= $%d", argPos))
		filterArgs = append(filterArgs, req.HeightFrom)
		argPos++
	}
	if req.HeightTo > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("height <= $%d", argPos))
		filterArgs = append(filterArgs, req.HeightTo)
		argPos++
	}
	if req.WeightFrom > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("weight >= $%d", argPos))
		filterArgs = append(filterArgs, req.WeightFrom)
		argPos++
	}
	if req.WeightTo > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("weight <= $%d", argPos))
		filterArgs = append(filterArgs, req.WeightTo)
		argPos++
	}

	// Базовый запрос
	baseQuery := `
        SELECT id, name, last_name, avatar_url
        FROM profiles
    `

	// Строим запрос для получения данных
	dataQuery := baseQuery
	if len(whereClauses) > 0 {
		dataQuery += " WHERE " + strings.Join(whereClauses, " AND ")
	}
	dataQuery += " ORDER BY id"

	// Добавляем пагинацию
	dataQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)

	// Объединяем все аргументы
	allArgs := append(filterArgs, req.Limit, offset)

	// 1. Получаем данные с пагинацией
	rows, err := s.db.QueryContext(ctx, dataQuery, allArgs...)
	if err != nil {
		return models.GetCatalogResponse{}, fmt.Errorf("failed to query items: %w", err)
	}
	defer rows.Close()

	// 2. Собираем результаты
	var items []models.CatalogItem
	for rows.Next() {
		var item models.CatalogItem
		if err := rows.Scan(
			&item.UserId, &item.Name, &item.LastName,
			&item.AvatarUrl,
		); err != nil {
			return models.GetCatalogResponse{}, fmt.Errorf("failed to scan item: %w", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return models.GetCatalogResponse{}, fmt.Errorf("rows error: %w", err)
	}

	// 3. Получаем общее количество элементов (только если есть фильтры)
	var totalItems int32
	if len(filterArgs) > 0 {
		countQuery := "SELECT COUNT(*) FROM profiles WHERE " + strings.Join(whereClauses, " AND ")
		err = s.db.QueryRowContext(ctx, countQuery, filterArgs...).Scan(&totalItems)
		if err != nil {
			return models.GetCatalogResponse{}, fmt.Errorf("failed to count items: %w", err)
		}
	} else {
		// Если нет фильтров, можно использовать более быстрый способ
		err = s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM profiles").Scan(&totalItems)
		if err != nil {
			return models.GetCatalogResponse{}, fmt.Errorf("failed to count items: %w", err)
		}
	}

	// 4. Рассчитываем общее количество страниц
	totalPages := totalItems / req.Limit
	if totalItems%req.Limit != 0 {
		totalPages++
	}

	return models.GetCatalogResponse{
		Items:      items,
		TotalPages: totalPages,
	}, nil
}

// getAvatarURL возвращает URL аватара на основе текущего аватара и списка фотографий
func getAvatarURL(currentAvatarURL string, photos []string) string {
    if len(photos) > 0 {
        return photos[0]  // Всегда используем первую фотографию как аватар
    }
    return ""  // Если фотографий нет, очищаем аватар
}
