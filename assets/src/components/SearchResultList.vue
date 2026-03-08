<script setup lang="ts">
import type { SearchResult } from '../types/search'
import ResultCard from './ResultCard.vue'

defineProps<{ results: SearchResult[] }>()
</script>

<template>
  <div v-if="results.length > 0">
    <ResultCard v-for="result in results" :key="result.url" :result="result" />
  </div>
</template>