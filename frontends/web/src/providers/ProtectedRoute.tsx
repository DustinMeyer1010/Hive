import { Navigate } from "react-router-dom"
import { useAuth } from "./authentication";
import type { ReactNode } from "react";


const ProtectedRoute: React.FC<{ children: ReactNode }> = ({children}) => {
    const { isLoggedIn } = useAuth()

    console.log(isLoggedIn)


    if (!isLoggedIn) return <Navigate to="/login" replace/>;
     return <>{children}</>
}

export default ProtectedRoute;