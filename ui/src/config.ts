

const Config = {
    projectName: "Opsie",
    version: import.meta.env.VITE_APP_VERSION,
    apiBaseUrl: "/api/v1",
    isDev: import.meta.env.VITE_APP_ENV === "development",
    isProd: import.meta.env.VITE_APP_ENV === "production",
} as const;

export default Config; 