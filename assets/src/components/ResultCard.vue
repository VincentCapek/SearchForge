<script setup lang="ts">
import type { SearchResult } from '../types/search'
defineProps<{ result: SearchResult }>()
</script>

<template>
  <article class="result-card">
    <div class="result-domain-row">
      <span>{{ result.domain }}</span>
    </div>

    <h2 class="text-xl leading-snug mb-0.5">
      <a
        :href="result.url"
        target="_blank"
        rel="noreferrer noopener"
        class="result-title-link"
      >{{ result.title }}</a>
    </h2>

    <span class="result-url-text">{{ result.url }}</span>
    <p class="result-snippet">{{ result.snippet }}</p>
  </article>
</template>
