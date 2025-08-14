import React, { createContext, useContext, useState, type ReactNode } from "react";

type Auth = {
    isLoggedIn: boolean;
    login: () => void;
    logout: () => void;
}

const AuthContext = createContext<Auth | undefined>(undefined)

export const AuthProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [isLoggedIn, setIsLoggedIn] = useState<boolean>(true);


    const login = () => setIsLoggedIn(true);
    const logout = () => setIsLoggedIn(false);


    return (
        <AuthContext.Provider value={{ isLoggedIn, login, logout}}>
        {children}
        </AuthContext.Provider>
    )
}

export function useAuth() {
    const context = useContext(AuthContext)
    if (!context) throw new Error("useAuth must be used within an AuthProvider")
    return context;
}