<template>
  <div class="form-group">
    <label v-if="label" :for="id">{{ label }}</label>
    <input
      v-if="type === 'email'"
      autocomplete="username"
      :id="id"
      :type="type"
      :placeholder="placeholder"
      :required="required"
      v-model="inputValue"
      @input="validateInput"
    />
    <input
      v-else
      :id="id"
      :type="type"
      :placeholder="placeholder"
      :required="required"
      :min="min"
      :max="max"
      :class="inputClass"
      v-model="inputValue"
      @input="validateInput"
    />
    <FieldErrorMsg :error="error" />
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  id: String,
  label: String,
  type: {
    type: String,
    default: 'text',
  },
  modelValue: String,
  placeholder: String,
  error: String,
  min: Number,
  max: Number,
  required: Boolean,
  inputClass: String,
})

const emit = defineEmits(['update:modelValue', 'update:error'])

const inputValue = ref(props.modelValue)

const validateInput = () => {
  let error = ''
  if (props.type === 'text') {
    const regex = /^[A-Za-zА-Яа-я\s]+$/
    if (!regex.test(inputValue.value)) {
      error = 'Имя должно содержать только буквы'
    }
  } else if (props.type === 'email') {
    const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    if (!regex.test(inputValue.value)) {
      error = 'Введите корректный email'
    }
  }
  emit('update:error', error)
}

watch(inputValue, (newValue) => {
  emit('update:modelValue', newValue)
})

watch(
  props,
  (newValue) => {
    inputValue.value = newValue.modelValue
  },
  { deep: true }
)
</script>

<style scoped>
.form-group {
  display: flex;
  position: relative;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  gap: 20px;
}

label {
  display: block;
  font-size: 15px;
  font-weight: 600;
}

input,
select {
  font-size: 15px;
  font-weight: 600;
  border: 1px solid #fff;
  background-color: #070707;
  padding: 8px;
  color: #fff;
  max-width: 390px;
  width: 100%;
}

select option {
  font-size: 13px;
  color: #f8f8f8;
  background-color: #070707;
  font-weight: 600;
}

input,
select:focus {
  outline: none;
}

input[type='number']::-webkit-outer-spin-button,
input[type='number']::-webkit-inner-spin-button {
  -webkit-appearance: none;
}

input[type='number'],
input[type='number']:hover,
input[type='number']:focus {
  appearance: none;
  -moz-appearance: textfield;
}

.half-number-input {
  width: 195px;
  text-align: center;
}

.third-number-input {
  width: 130px;
  text-align: center;
}

@media (max-width: 918px) {
  input,
  select {
    width: 230px;
  }
}
</style>
