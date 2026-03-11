<script setup lang="ts">
import type { SearchEngine } from '../composables/useSearch'

defineProps<{ loading: boolean }>()
const model = defineModel<string>()
const engine = defineModel<SearchEngine>('engine')
const emit = defineEmits<{ search: [] }>()
</script>

<template>
  <div class="flex flex-col gap-3">
    <!-- Engine selector -->
    <div class="flex gap-1.5 p-1 bg-[#f1f3f4] rounded-lg w-fit">
      <button
        type="button"
        :class="[
          'px-3 py-1.5 text-xs font-medium rounded-md cursor-pointer transition-colors border-0',
          engine === 'ddg'
            ? 'bg-surface-base text-text-primary shadow-sm'
            : 'bg-transparent text-text-muted hover:text-text-primary',
          loading && 'cursor-not-allowed opacity-60'
        ]"
        :disabled="loading"
        @click="engine = 'ddg'"
      >
        DuckDuckGo
      </button>
      <button
        type="button"
        :class="[
          'px-3 py-1.5 text-xs font-medium rounded-md cursor-pointer transition-colors border-0',
          engine === 'brave'
            ? 'bg-surface-base text-text-primary shadow-sm'
            : 'bg-transparent text-text-muted hover:text-text-primary',
          loading && 'cursor-not-allowed opacity-60'
        ]"
        :disabled="loading"
        @click="engine = 'brave'"
      >
        Brave
      </button>
    </div>

    <!-- Search form -->
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
  </div>
</template>
