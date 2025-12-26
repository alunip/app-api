import { useState, useEffect } from 'react';
import type { AppConfig } from '../models/config';
import { getConfig } from '../services/configService';

/**
 * Return type for useConfig hook
 */
export interface UseConfigReturn {
  data: AppConfig | null;
  loading: boolean;
  error: Error | null;
  refetch: () => Promise<void>;
}

/**
 * Custom hook for fetching and managing application configuration
 * Automatically fetches config on mount and provides manual refetch capability
 *
 * @returns Object containing data, loading state, error state, and refetch function
 *
 * @example
 * const { data, loading, error, refetch } = useConfig();
 *
 * if (loading) return <div>Loading...</div>;
 * if (error) return <div>Error: {error.message}</div>;
 * if (!data) return null;
 *
 * return <div>{data.name} - v{data.version}</div>;
 */
export const useConfig = (): UseConfigReturn => {
  const [data, setData] = useState<AppConfig | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<Error | null>(null);

  const fetchConfig = async () => {
    try {
      setLoading(true);
      setError(null);
      const config = await getConfig();
      setData(config);
    } catch (err) {
      setError(err instanceof Error ? err : new Error('Unknown error occurred'));
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchConfig();
  }, []);

  const refetch = async () => {
    await fetchConfig();
  };

  return { data, loading, error, refetch };
};
