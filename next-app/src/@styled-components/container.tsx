import {Container as C} from "@mui/material"
import React, { ReactNode } from "react"
export default function Container ({children, className = ""}: {children: ReactNode, className?: string}) {
        return <C className={`${className}`}>
           {children} 
        </C>
}
