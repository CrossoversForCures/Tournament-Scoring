import { env as privateEnv } from '$env/dynamic/private';
import { dev } from '$app/environment';

interface Env {
    ADMIN_USERNAME: string;
    ADMIN_PASSWORD: string;
    // Add other variables as needed
}

// Create a type-safe environment object
const env: Env = {
    ADMIN_USERNAME: privateEnv.VITE_ADMIN_USERNAME || '',
    ADMIN_PASSWORD: privateEnv.VITE_ADMIN_PASSWORD || '',
};

// In development, verify all variables are set
if (dev) {
    const unsetEnv = Object.entries(env).filter(([_, v]) => v === '');
    if (unsetEnv.length > 0) {
        throw new Error(`Missing environment variables: ${unsetEnv.map(([k]) => k).join(', ')}`);
    }
}

export default env;