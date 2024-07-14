import {FormEvent, useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {DoRequest} from "../wailsjs/go/main/App";
import { http_types } from '../wailsjs/go/models';

function App() {
    const [result, setResult] = useState<http_types.Response | undefined>();

    function submit(event: FormEvent): void {
        setResult(undefined)
        event.preventDefault();
        const data = new FormData(event.target as HTMLFormElement);
        console.log(data.get('method'), data.get('url')); 
        DoRequest(data.get('url') as string, data.get('method') as string).then(response => {
            setResult(response);
        })
    }
    return (
        <div id="App">
            <form onSubmit={submit}>
                <select name="method" id="method">
                    <option value="GET">GET</option>
                    <option value="POST">POST</option>
                    <option value="PUT">PUT</option>
                    <option value="DELETE">DELETE</option>
                </select>
                <input type="text" name='url'/>
                <button type='submit'>Send</button>
            </form>
            <section>
                {result && <article>
                    <p>Status</p>
                    <p>{result.status}</p>
                    <p>Response</p>
                    <p>{result.body}</p>
                </article>}
            </section>
        </div>
    )
}

export default App
