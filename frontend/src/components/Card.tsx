import { useState } from 'react'


function Card(){

    const [count, setcount] = useState(0)

    function handleClick(){
        return setcount(count + 1)
    }

    return (
        <div className="card" onClick={handleClick}>
            <h1>{count}</h1>
        </div>
    )
}

export default Card