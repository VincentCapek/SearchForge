<script setup lang="ts">
defineProps<{ loading: boolean }>()
const model = defineModel<string>()
const emit = defineEmits<{ search: [] }>()
</script>

<template>
  <form class="flex gap-2.5 items-center" @submit.prevent="emit('search')">
    <input
      v-model="model"
      type="text"
      placeholder="Search anything..."
      autocomplete="off"
      class="search-input"
      :disabled="loading"
    />
    <button type="submit" class="btn-search" :disabled="loading">
      {{ loading ? 'Searching…' : 'Search' }}
    </button>
  </form>
</template>
