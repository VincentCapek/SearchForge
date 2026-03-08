import { ref } from 'vue'
import type { SearchResult } from '../types/search'
import { searchDuckDuckGo } from '../lib/api'

export function useSearch() {
  const query = ref('')
  const loading = ref(false)
  const error = ref('')
  const results = ref<SearchResult[]>([])
  const hasSearched = ref(false)

  async function search() {
    error.value = ''
    loading.value = true
    hasSearched.value = true

    try {
      results.value = await searchDuckDuckGo(query.value)
    } catch (err) {
      results.value = []
      error.value = err instanceof Error ? err.message : 'Unknown error'
    } finally {
      loading.value = false
    }
  }

  return { query, loading, error, results, hasSearched, search }
}
