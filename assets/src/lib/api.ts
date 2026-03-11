import type { SearchResult } from "../types/search";

export async function searchDuckDuckGo(query: string): Promise<SearchResult[]> {
    const trimmedQuery = query.trim();
    if (trimmedQuery === "") {
        return [];
    }

    const response = await fetch(`/api/search/ddg?q=${encodeURIComponent(trimmedQuery)}`);
    if (!response.ok) {
        let message = `HTTP ${response.status}`

        try {
            const errorData = await response.json();
            message = errorData.error || message;
        } catch (error) {
            // ignore JSON parse errors
        }

        throw new Error(message);
    }

    return response.json() as Promise<SearchResult[]>;
}

export async function searchBrave(query: string): Promise<SearchResult[]> {
    const trimmedQuery = query.trim();
    if (trimmedQuery === "") {
        return [];
    }

    const response = await fetch(`/api/search/brave?q=${encodeURIComponent(trimmedQuery)}`);
    if (!response.ok) {
        let message = `HTTP ${response.status}`

        try {
            const errorData = await response.json();
            message = errorData.error || message;
        } catch (error) {
            // ignore JSON parse errors
        }

        throw new Error(message);
    }

    return response.json() as Promise<SearchResult[]>;
}