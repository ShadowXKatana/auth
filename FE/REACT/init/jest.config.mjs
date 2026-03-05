/** @type {import('jest').Config} */
const config = {
    testEnvironment: "jest-environment-jsdom",
    setupFilesAfterEnv: ["<rootDir>/jest.setup.ts"],
    transform: {
        "^.+\\.tsx?$": [
            "ts-jest",
            {
                tsconfig: "tsconfig.app.json",
            },
        ],
    },
    moduleNameMapper: {
        "^@/(.*)$": "<rootDir>/src/$1",
        "\\.(css|less|scss|sass)$": "identity-obj-proxy",
        "\\.(gif|ttf|eot|svg|png|jpg|jpeg|webp)$": "<rootDir>/__mocks__/fileMock.js",
    },
};

export default config;
