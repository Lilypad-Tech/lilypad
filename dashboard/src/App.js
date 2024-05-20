import React, { useEffect, useState } from 'react';
import io from 'socket.io-client';

// import './globals.css';
// import './output.css';




// import { json } from 'stream/consumers';
const ENDPOINT = 'http://localhost:8000'; // Change this to your server's endpoint

const App = () => {
  const [logupdates, setLogUpdates] = useState([]);
  const [jobupdates, setJobUpdates] = useState([]);
  const [response, setResponse] = useState('');
  const [stdout, setStdout] = useState('');
  const [matcher, setMatcher] = useState('');
  const [deal, setDeal] = useState('');
  const [result, setResult] = useState('');
  const [imgresult, setImgresult] = useState('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAA1JREFUGFdjOHPmzH8ACDADZKt3GNsAAAAASUVORK5CYII=');
  const [selectedValue, setSelectedValue] = useState("github.com/arsen3d/lilypad-module-lilysay:0.1.0 Message=");
  const [promptValue, setPromptValue] = useState("A happy little tree");
  const [imageDisplay, setImageDisplay] = useState("none");

  
  const handleSelectChange = (event) => {
    setSelectedValue(event.target.value);
  };
  const handleChange = (event) => {
    setPromptValue(event.target.value);
  }

  //{transports: ['websocket']}
  const socket = io();
  // let socket;
  function deepSearch(obj, targetValue, prop) {
    for (const key in obj) {
      if (typeof obj[key] === 'object') {
        const result = deepSearch(obj[key], targetValue, prop);
        if (result !== null) {
          return result;
        }
      } else if (obj[key] === targetValue) {
        return { key, value: obj[prop] };
      }
    }
    return null;
  }
  useEffect(() => {
    // const str = window.location.origin.replace("http","ws")
    // socket = io();//'http://172.23.16.133:8080');
    // socket.connect()
    // socket.on('updates', (data) => {
    //   console.log(data)
    //     // setUpdates((prevUpdates) => [...prevUpdates, data]);
    // });

    // Fetch pending updates when component mounts
    fetchLogUpdates();
    fetchJobUpdates();
    // return () => {
    //   socket.disconnect();
    // };


    socket.on('connect', () => {
      console.log('Connected to server');
    });


    socket.on('deal', data => {
      console.log('Received deal:', data);
      setDeal(data)
      fetchJobUpdates()
      //setResponse(data);
    });

    socket.on('matcher', data => {
      console.log('Received matcher:', data);
      setMatcher(data)
      //setResponse(data);
    });
    socket.on('result', data => {
      console.log('result deal:', data);

      //setResult(data)
      //setResponse(data);
    });
    socket.on('imgresult', data => {
      console.log('result deal:', data);
      // setImgresult(data)
      //setResponse(data);
    });


    socket.on('reply', data => {
      //console.log('Received reply:', data);
      setStdout(data)
      //setResponse(data);
    });
    socket.on('update', data => {
      console.log('Received reply:', data);
      setLogUpdates(data);
    });

    socket.on('bye', data => {
      console.log('Received bye:', data);
      socket.disconnect(); // Disconnect from the server when receiving 'bye'
    });
    socket.on('resultfolder', data => {
      console.log('resultfolder:', data);
      try {
        const llama_response = JSON.parse(JSON.parse(data).result.Details.contents[3].contents[0].contents[0].content).response
        setResult(llama_response)
      } catch (e) {
        setResult(deepSearch(JSON.parse(data), 'stdout', "content").value)
      }
      if (deepSearch(JSON.parse(data), 'image/png', "relpath") != null) {
        setImgresult("/files" + deepSearch(JSON.parse(data), 'image/png', "relpath").value)
        setImageDisplay("visible")
      }else{
        setImageDisplay("none")
      }
      //JSON.parse(JSON.parse(data).result.Details.contents[3].contents[0].contents[0].content).response
      // console.log(deepSearch(JSON.parse(data), 'stdout', "content"));
      //const stdoutContentBase64 = JSON.parse(data).Details.contents.find(item => item.name === 'stdout').content;

      // Decoding the Base64 content
      //const stdoutContentDecoded = atob(deepSearch(JSON.parse(data), 'stdout',"content").value);

      // console.log("img","/files"+deepSearch(JSON.parse(data), 'image/png',"relpath").value);




      //console.log("imagepath",deepSearch(JSON.parse(data), 'image/png',"path"))
      // socket.disconnect(); // Disconnect from the server when receiving 'bye'
    });

    return () => socket.disconnect(); // Clean up on unmount
  }, []);



  const fetchLogUpdates = async () => {
    try {
      const response = await fetch(window.location.origin + "/log-updates");
      const data = await response.json();
      console.log(data)
      setLogUpdates(data);
    } catch (error) {
      console.error('Error fetching pending updates:', error);
    }
  };
  const fetchJobUpdates = async () => {
    try {
      const response = await fetch(window.location.origin + "/job-updates");
      const data = await response.json();
      console.log(data)
      setJobUpdates(data);
    } catch (error) {
      console.error('Error fetching pending updates:', error);
    }
  };

  const pingServer = async () => {
    //   socket = io(window.location.origin);//'http://172.23.16.133:8080');
    socket.emit('message', "messageData");
    socket.emit("notice", "Starting");
  }
  const genImage = async () => {
    //   socket = io(window.location.origin);//'http://172.23.16.133:8080');
    // socket.emit('message', "messageData");
    socket.emit("task", "genImage");
  }
  const runModule = async () => {
    //   socket = io(window.location.origin);//'http://172.23.16.133:8080');
    // socket.emit('message', "messageData");
    //runsdxl sdxl:v0.9-lilypad1 PROMPT='moo'
    console.log("task", selectedValue.split(" ")[0] + " -i " + selectedValue.split(" ")[1] + `${promptValue}`)
    socket.emit("task", selectedValue.split(" ")[0] + " -i " + selectedValue.split(" ")[1] + `${promptValue}`);
    //-i PromptEnv='PROMPT=big monkey'
  }

  return (


    <div className="bg-zinc-950 dark:bg-white" >


      {/* <h1 class="text-3xl font-bold underline">
       Monitor
      </h1> */}
      {/* <iframe src="http://localhost:3000/dashboard-solo/new?orgId=1&refresh=auto&from=1715712433852&to=1715714233852&panelId=1" width="450" height="200" frameborder="0"></iframe> */}
      <select value={selectedValue} onChange={handleSelectChange}>
        <option value="github.com/arsen3d/lilypad-module-lilysay:0.1.0 Message=">github.com/arsen3d/lilypad-module-lilysay:0.1.0</option>
        <option value="github.com/Lilypad-Tech/lilypad-module-sdxl:v0.9-lilypad1 PromptEnv=PROMPT=">github.com/Lilypad-Tech/lilypad-module-sdxl:v0.9-lilypad1</option>
        <option value="github.com/arsen3d/lilypad-module-ollama-pipeline:test6 .prompt=">github.com/arsen3d/lilypad-module-ollama-pipeline:test5</option>
        <option value="github.com/Lilypad-Tech/lilypad-module-ollama-pipeline:llama3-8b-lilypad1 prompt=">github.com/Lilypad-Tech/lilypad-module-ollama-pipeline:llama3-8b-lilypad1</option>
        {/* https://github.com/Lilypad-Tech/lilypad-module-ollama-pipeline/tree/llama3-8b-lilypad1 */}
        {/* https://github.com/Lilypad-Tech/lilypad-module-ollama-pipeline/tree/llama3-8b-lilypad1 */}
      </select>
      <input type="text" value={promptValue} onChange={handleChange} />
      <button onClick={runModule}>Run Module</button>
      {/* <Button asChild>
        <Link href="/login">Login</Link>
      </Button> */}
      {/* <br/>
            <button onClick={pingServer}>Cow Say</button>
            <button onClick={genImage}>SDXL Image Gen</button> */}
      <br />
      Matches:
      <br />
      <code>{matcher}</code>
      <br />
      Deal Status:
      <br />
      <code>{deal}</code>
      <br />
      Result:
      <br />
      <div style={{ display: 'flex', justifyContent: 'space-between' }}>
        <pre>{result}</pre>
        <img style={{ display: imageDisplay }} width={256} height={256} src={imgresult} />
      </div>
      <br />
      Messages:
      <br />
      <code>{stdout}</code>
      <h1>Job Metrics</h1>

      <table border={1}>
        <tr><td>Start</td><td>Updated</td><td>Duration</td><td>Module</td><td>Status</td></tr>
        {(jobupdates != null && jobupdates.length) ? jobupdates.map((update) => (

          <tr key={update.id}>
            <td>{new Date(update.time_start).toLocaleString('en-US')}</td>
            <td>{new Date(update.time_update).toLocaleString('en-US')}  </td>
            <td>{new Date((new Date(update.time_update) - new Date(update.time_start))).getSeconds()} Seconds</td>
            <td>{update.module_id} </td>
            <td>{update.status}</td>

          </tr>

        )) : null}
      </table>
      <h1>Logs</h1>

      <ul>
        {logupdates.map((update) => (
          <li key={update.id}> {new Date(update.timestamp).toLocaleString('en-US')}  {update.message}</li>
        ))}
      </ul>
    </div>
  );
};

export default App;
