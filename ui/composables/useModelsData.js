import { ref } from 'vue'
import modelImage from '~/assets/test-photo.jpg'

export const useModelsData = () => {
  const models = ref([
    {
      id: 1,
      name: 'Анна Иванова',
      image: modelImage,
      height: 175,
      weight: 50,
      parameters: '90-60-90',
      age: 25,
      experience: '5 лет',
      description:
        'Профессиональная модель с опытом работы на международных показах. Участница недели моды в Милане и Париже.',
      portfolio: [
        { id: 1, image: modelImage, title: 'Фотосессия для Vogue' },
        { id: 2, image: modelImage, title: 'Реклама парфюмерии' },
        { id: 3, image: modelImage, title: 'Показ Весна/Лето 2024' },
      ],
    },
    {
      id: 2,
      name: 'Мария Петрова',
      image: modelImage,
      height: 178,
      weight: 50,
      parameters: '86-58-88',
      age: 23,
      experience: '3 года',
      description:
        'Специализация на рекламных съемках. Сотрудничала с ведущими брендами косметики и одежды.',
      portfolio: [
        { id: 1, image: modelImage, title: 'Реклама спортивной одежды' },
        { id: 2, image: modelImage, title: "Кампания для L'Oreal" },
      ],
    },
    {
      id: 3,
      name: 'Елена Сидорова',
      image: modelImage,
      height: 176,
      weight: 50,
      parameters: '88-61-92',
      age: 24,
      experience: '4 года',
      description:
        'Модель для печатных изданий и каталогов. Опыт работы с фотографами из Европы и США.',
      portfolio: [
        { id: 1, image: modelImage, title: 'Каталог модной одежды' },
        { id: 2, image: modelImage, title: 'Съемка для журнала Elle' },
      ],
    },
    {
      id: 4,
      name: 'Ольга Кузнецова',
      image: modelImage,
      height: 180,
      weight: 50,
      parameters: '92-62-92',
      age: 26,
      experience: '6 лет',
      description:
        'Модель и актриса. Снималась в рекламных роликах и фильмах. Опыт работы в театре.',
      portfolio: [
        { id: 1, image: modelImage, title: 'Съемки в рекламе автомобиля' },
        { id: 2, image: modelImage, title: 'Эпизодическая роль в сериале' },
      ],
    },
    {
      id: 5,
      name: 'Ирина Смирнова',
      image: modelImage,
      height: 174,
      weight: 50,
      parameters: '86-59-89',
      age: 22,
      experience: '2 года',
      description:
        'Начинающая модель с потенциалом. Победительница региональных конкурсов красоты.',
      portfolio: [
        { id: 1, image: modelImage, title: 'Конкурс Мисс Весна 2023' },
        { id: 2, image: modelImage, title: 'Фотосессия для местного журнала' },
      ],
    },
    {
      id: 6,
      name: 'Татьяна Васильева',
      image: modelImage,
      height: 177,
      weight: 50,
      parameters: '89-60-91',
      age: 25,
      experience: '4 года',
      description:
        'Опытная модель для показов и фотосессий. Работала с известными дизайнерами.',
      portfolio: [
        { id: 1, image: modelImage, title: 'Показ одежды молодых дизайнеров' },
        { id: 2, image: modelImage, title: 'Фотосессия для каталога' },
      ],
    },
    {
      id: 7,
      name: 'Екатерина Козлова',
      image: modelImage,
      height: 179,
      weight: 50,
      parameters: '91-61-92',
      age: 27,
      experience: '7 лет',
      description:
        'Профессиональная модель высокого уровня. Международный опыт работы, контракты с ведущими агентствами.',
      portfolio: [
        { id: 1, image: modelImage, title: 'Показ в Нью-Йорке' },
        { id: 2, image: modelImage, title: 'Съемка рекламы для Dior' },
      ],
    },
    {
      id: 8,
      name: 'Наталья Морозова',
      image: modelImage,
      height: 175,
      weight: 50,
      parameters: '85-58-87',
      age: 23,
      experience: '3 года',
      description:
        'Стильная модель для фотосессий и каталогов. Особый талант к позированию.',
      portfolio: [
        { id: 1, image: modelImage, title: 'Фотосессия в студии' },
        { id: 2, image: modelImage, title: 'Каталог аксессуаров' },
      ],
    },
    {
      id: 9,
      name: 'Алина Белова',
      image: modelImage,
      height: 176,
      weight: 50,
      parameters: '87-59-88',
      age: 24,
      experience: '4 года',
      description:
        'Модель и визажист. Понимает процесс съемки с обеих сторон камеры.',
      portfolio: [
        { id: 1, image: modelImage, title: 'Реклама косметики' },
        { id: 2, image: modelImage, title: 'Мастер-класс по макияжу' },
      ],
    },
    {
      id: 10,
      name: 'Ксения Волкова',
      image: modelImage,
      height: 178,
      weight: 50,
      parameters: '90-62-93',
      age: 26,
      experience: '5 лет',
      description:
        'Специализация на фешн-фотографии и показах. Сотрудничает с известными фотографами.',
      portfolio: [
        { id: 1, image: modelImage, title: 'Показ вечерних платьев' },
        { id: 2, image: modelImage, title: 'Съемка для глянцевого журнала' },
      ],
    },
    {
      id: 11,
      name: 'Дарья Соколова',
      image: modelImage,
      height: 177,
      weight: 50,
      parameters: '88-60-90',
      age: 25,
      experience: '5 лет',
      description:
        'Разносторонняя модель. Работа с дизайнерами одежды и ювелирных украшений.',
      portfolio: [
        { id: 1, image: modelImage, title: 'Показ ювелирных украшений' },
        { id: 2, image: modelImage, title: 'Реклама вечерних платьев' },
      ],
    },
    {
      id: 12,
      name: 'Полина Лебедева',
      image: modelImage,
      height: 174,
      weight: 50,
      parameters: '85-58-86',
      age: 22,
      experience: '2 года',
      description: 'Молодая перспективная модель. Энергичная и креативная.',
      portfolio: [
        { id: 1, image: modelImage, title: 'Уличная фотосессия' },
        { id: 2, image: modelImage, title: 'Молодежная реклама' },
      ],
    },
  ])

  const getModelById = (id) => {
    const modelId = parseInt(id)
    return models.value.find((model) => model.id === modelId) || null
  }

  return {
    models,
    getModelById,
  }
}
