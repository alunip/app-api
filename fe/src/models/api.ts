/**
 * Generic API response wrapper used by all endpoints
 */
export interface ApiResponse<T> {
  data?: T;
  error?: string;
  timestamp: string;
}
