<script setup lang="ts">
import SearchBar from './components/SearchBar.vue'
import SearchResultList from './components/SearchResultList.vue'
import { useSearch } from './composables/useSearch'

const { query, loading, error, results, hasSearched, search } = useSearch()
</script>

<template>
  <div class="page-shell">
    <div class="page-container">

      <!-- Header -->
      <header class="mb-8">
        <p class="text-xs font-semibold tracking-[0.12em] uppercase text-[#9aa0a6] mb-1.5">SerpGo</p>
        <h1 class="text-[1.75rem] font-semibold text-[#202124] leading-tight m-0">Search the web</h1>
      </header>

      <!-- Search bar -->
      <SearchBar v-model="query" :loading="loading" @search="search" />

      <!-- States -->
      <div class="mt-7">

        <!-- Loading -->
        <div v-if="loading" class="feedback-panel">
          <span class="w-4 h-4 rounded-full border-2 border-[#e5e7eb] border-t-[#4285f4] animate-spin shrink-0" />
          <span>Searching…</span>
        </div>

        <!-- Error -->
        <p v-else-if="error" class="error-panel">{{ error }}</p>

        <!-- Results -->
        <template v-else-if="hasSearched">
          <p class="count-line">
            {{ results.length }} result{{ results.length !== 1 ? 's' : '' }}
          </p>
          <p v-if="results.length === 0" class="feedback-panel">No results found for your query.</p>
          <SearchResultList v-else :results="results" />
        </template>

        <!-- Initial state -->
        <p v-else class="feedback-panel">Enter a search term above to get started.</p>

      </div>
    </div>
  </div>
</template>
