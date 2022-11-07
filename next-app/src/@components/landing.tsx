import { TextField } from "@mui/material";
import Container from "../@styled-components/container";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSearch } from '@fortawesome/free-solid-svg-icons'
import { Dispatch, SetStateAction, useEffect, useState } from "react";

const DropDown = ({input, setInput}: {input: string, setInput: Dispatch<SetStateAction<string>>}) => {
    useEffect(() => {
        console.log(input)
    }, [input])
    return <div className="bg-gray-100 w-full rounded-lg mx-4 h-1/2">
        
    </div>
}

export default function Landing () {
    const [focused, setFocused] = useState(true)
    const [input, setInput] = useState("")

    return (<>
       <div className="bg-black h-screen w-full center">
            <Container className={`m-0 p-4 h-full w-full bg-[#17181a] flex flex-col ${focused ? "justify-start" : "justify-end"} items-center py-12 rounded-3xl`}>
                {!focused && <div className="">
                    <p className="text-primary-dark text-6xl mb-4"> GCT File Server </p>
                    <p className="text-gray-300 border-b-2 pb-4 border-b-gray-300 text-lg text-justify">Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit</p>
                </div>}
                <div className="m-4 w-full">
                    <p className="uppercase text-primary-light text-lg my-4"> search: </p>
                    <TextField onChange={(e) => setInput(e.target.value)} value={input} onFocus={() => setFocused(true)} onBlur={() => setFocused(true)} placeholder="Type ..." className="bg-white focus:rounded-lg w-full outline-none rounded-lg text-black" InputProps={{
                        endAdornment: <FontAwesomeIcon icon={faSearch}/> 
                    }}></TextField>
                </div>
                {focused && <DropDown input={input} setInput={setInput}/>}
                {!focused && <div className="mt-12 h-1/4 mb-4">
                    <p className="text-gray-300 my-4"> Quick searches </p>    
                    <div className="w-full h-full bg-primary-dark rounded-lg">

                    </div>
                </div>}
            </Container>
                </div>
    </>)
}
