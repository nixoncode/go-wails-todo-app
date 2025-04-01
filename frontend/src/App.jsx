import {useState} from 'react';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e) => setName(e.target.value);
    const updateResultText = (result) => setResultText(result);
    

    function greet() {
        Greet(name).then(updateResultText);
    }
    const [one, setOne] = useState(0);
    const [two, setTwo] = useState(0);
    const updateOne = (e) => setOne(e.target.value);
    const updateTwo = (e) => setTwo(e.target.value);
    const [sum, setSum] = useState("");
    function add() {
        const sum = parseInt(one) + parseInt(two);
        setSum(`The sum of ${one} and ${two} is ${sum}`);
    }

    return (
        <div id="App">
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
                <button className="btn" onClick={greet}>Greet</button>
            </div>
            <hr />
            <input type="number" name="one" onChange={updateOne} /> 
            + 
            <input type="number" name="one" onChange={updateTwo} />
            <button className="btn" onClick={add}>Add</button>
            <div id="result" className="result">{sum}</div>
            <hr />

        </div>
    )
}

export default App
