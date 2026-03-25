import api from './api';
import type { ApiResponse } from '../models/api';
import type { AppConfig } from '../models/config';

/**
 * Fetches application configuration from the backend
 * @returns Promise resolving to AppConfig
 * @throws Error if request fails or response is invalid
 */
export const getConfig = async (): Promise<AppConfig> => {
  try {
    const response = await api.get<ApiResponse<AppConfig>>('/config');

    if (response.data.error) {
      throw new Error(response.data.error);
    }

    if (!response.data.data) {
      throw new Error('Configuration data not found in response');
    }

    return response.data.data;
  } catch (error) {
    if (error instanceof Error) {
      throw new Error(`Failed to fetch configuration: ${error.message}`);
    }
    throw new Error('Failed to fetch configuration: Unknown error');
  }
};
