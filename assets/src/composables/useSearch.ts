import { ref } from 'vue'
import type { SearchResult } from '../types/search'
import { searchDuckDuckGo, searchBrave } from '../lib/api'

export type SearchEngine = 'ddg' | 'brave'

export function useSearch() {
  const query = ref('')
  const selectedEngine = ref<SearchEngine>('ddg')
  const loading = ref(false)
  const error = ref('')
  const results = ref<SearchResult[]>([])
  const hasSearched = ref(false)

  async function search() {
    error.value = ''
    loading.value = true
    hasSearched.value = true

    try {
      if (selectedEngine.value === 'brave') {
        results.value = await searchBrave(query.value)
      } else {
        results.value = await searchDuckDuckGo(query.value)
      }
    } catch (err) {
      results.value = []
      error.value = err instanceof Error ? err.message : 'Unknown error'
    } finally {
      loading.value = false
    }
  }

  return { query, selectedEngine, loading, error, results, hasSearched, search }
}
