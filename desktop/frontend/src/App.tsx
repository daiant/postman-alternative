import { FormEvent, useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { DoRequest } from "../wailsjs/go/main/App";
import { http_types } from '../wailsjs/go/models';

function App() {
    const [result, setResult] = useState<http_types.Response | undefined>();
    const [params, setParams] = useState(1);

    function submit(event: FormEvent): void {
        setResult(undefined)
        event.preventDefault();
        const data = new FormData(event.target as HTMLFormElement);
        const params = [data.getAll('param_name'), data.getAll('param_value')];
        console.log(params)
        DoRequest(data.get('url') as string, data.get('method') as string).then(response => {
            setResult(response);
        })
    }
    function range(length: number): Array<number> {
      return [...Array(length).keys() ].map((i) => i);
    }
    return (
        <div id="App">
            <form onSubmit={submit}>
                <fieldset>
                    <select name="method" id="method">
                        <option value="GET">GET</option>
                        <option value="POST">POST</option>
                        <option value="PUT">PUT</option>
                        <option value="DELETE">DELETE</option>
                    </select>
                    <input type="text" name='url' />
                    <button type='submit'>Send</button>
                </fieldset>
                <fieldset>
                  <legend>Params</legend>
                  {range(params).map(v => <div key={v}>
                    <input type="text" name='param_name' required />
                    <input type="text" name='param_value'required />
                  </div>)}
                  <button onClick={() => setParams(v => v+1)}>Afegir mes</button>
                </fieldset>
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
