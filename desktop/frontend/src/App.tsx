import { FormEvent, useEffect, useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { DoRequest, GetSavedRequestsInWorkspace } from "../wailsjs/go/main/App";
import { http_types } from '../wailsjs/go/models';

type RequestFile = string;
type RequestFiles = {
  [key in string]: RequestFiles | RequestFile;
}

function App() {
  const [result, setResult] = useState<http_types.Response | undefined>();
  const [params, setParams] = useState(0);
  const [headers, setHeaders] = useState(0);
  const [savedRequests, setSavedRequests] = useState<RequestFiles>({});

  useEffect(() => {
    GetSavedRequestsInWorkspace().then(a => setSavedRequests(a))
  })
  function submit(event: FormEvent): void {
    setResult(undefined)
    event.preventDefault();
    const data = new FormData(event.target as HTMLFormElement);
    const params = [data.getAll('param_name'), data.getAll('param_value')];
    const request = {
      url: data.get('url'),
      method: data.get('method'),
      params: [data.getAll('param_name'), data.getAll('param_value')],
      body: data.get('body')
    } as http_types.Request;

    DoRequest(request).then(response => {
      setResult(response);
    })
  }
  function range(length: number): Array<number> {
    return [...Array(length).keys()].map((i) => i);
  }
  function getFolderDetails(files: RequestFiles): React.ReactNode {
    return Object.entries(files).map(value => <div key={value[0]} style={{ paddingInlineStart: 16 }}>
        {/* <p>{JSON.stringify(value[1], null, 2)}</p> */}
        {typeof value[1] === 'object' && <details>
          <summary>{value[0]}</summary> 
          {getFolderDetails(value[1])}
          </details>}
        {typeof value[1] !== 'object' && <p role='button'>{value[1]}</p>}
        {/* {Array.isArray(value[1]) && value[1].map(rf => <p role='button' key={rf.name}>{rf.name}</p>)}

        {!Array.isArray(value[1]) && getDetails(value[1], padding = true)} */}
    </div>)
  }
  function getDetails(files: RequestFiles): React.ReactNode {
    return <details>
      <summary>Workspace</summary>
      {getFolderDetails(files)}
    </details>
  }
  return (
    <div id="App">
      <fieldset style={{ textAlign: 'left' }}>
        <legend>Requests</legend>
        {getDetails(savedRequests)}
      </fieldset>
      <form onSubmit={submit}>
        <fieldset>
          <select name="method" id="method">
            <option value="GET">GET</option>
            <option value="POST">POST</option>
            <option value="PUT">PUT</option>
            <option value="DELETE">DELETE</option>
          </select>
          <input type="text" name='url' defaultValue="http://httpbin.org/get" />
          <button type='submit'>Send</button>
        </fieldset>
        <fieldset>
          <legend>Params</legend>
          {range(params).map(v => <div key={v}>
            <input type="text" name='param_name' required />
            <input type="text" name='param_value' required />
          </div>)}
          <button onClick={() => setParams(v => v + 1)}>Afegir mes</button>
        </fieldset>
        <fieldset>
          <legend>Headers</legend>
          {range(headers).map(v => <div key={v}>
            <input type="text" name='header_name' required />
            <input type="text" name='header_value' required />
          </div>)}
          <button onClick={() => setHeaders(h => h + 1)}>Afegir mes</button>
        </fieldset>
        <fieldset>
          <legend>Body</legend>
          <textarea name="body" id=""></textarea>
        </fieldset>
      </form>
      <section>
        {result && <fieldset>
          <table style={{ textAlign: 'left', borderSpacing: 16 }}>
            <thead>
              <tr>
                <th>Status</th>
                <th>Response</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td style={{ verticalAlign: 'top' }}>{result.status}</td>
                <td style={{ verticalAlign: 'top', whiteSpace: 'break-spaces' }}>{result.body}</td>
              </tr>
            </tbody>
          </table>
        </fieldset>
        }
      </section>
    </div>
  )
}

export default App
