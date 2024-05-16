import React, { useEffect, useState } from 'react';
import io from 'socket.io-client';
// import { json } from 'stream/consumers';
const ENDPOINT = 'http://localhost:8000'; // Change this to your server's endpoint

const App = () => {
    const [updates, setUpdates] = useState([]);
    const [response, setResponse] = useState('');
    const [stdout, setStdout] = useState('');
    const [matcher, setMatcher] = useState('');
    const [deal, setDeal] = useState('');
    const [result, setResult] = useState('');
    const [imgresult, setImgresult] = useState('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAA1JREFUGFdjOHPmzH8ACDADZKt3GNsAAAAASUVORK5CYII=');
    const [selectedValue, setSelectedValue] = useState("github.com/Lilypad-Tech/lilypad-module-sdxl:v0.9-lilypad1 PromptEnv=PROMPT=");
    const [promptValue, setPromptValue] = useState("A happy little tree");
    const handleSelectChange = (event) => {
      setSelectedValue(event.target.value);
    };
    const handleChange  = (event) => {     
      setPromptValue(event.target.value); 
    }
    const socket = io({transports: ['websocket']});
    // let socket;
    function deepSearch(obj, targetValue,prop) {
      for (const key in obj) {
          if (typeof obj[key] === 'object') {
              const result = deepSearch(obj[key], targetValue,prop);
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
        fetchPendingUpdates();

        // return () => {
        //   socket.disconnect();
        // };
       

        socket.on('connect', () => {
          console.log('Connected to server');
        });
    
        
        socket.on('deal', data => {
          console.log('Received deal:', data);
          setDeal(data)
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
            setUpdates(data);
          });
    
        socket.on('bye', data => {
          console.log('Received bye:', data);
          socket.disconnect(); // Disconnect from the server when receiving 'bye'
        });
        socket.on('resultfolder', data => {
          console.log('resultfolder:', data);
          console.log(deepSearch(JSON.parse(data), 'stdout',"content"));
          //const stdoutContentBase64 = JSON.parse(data).Details.contents.find(item => item.name === 'stdout').content;

          // Decoding the Base64 content
          //const stdoutContentDecoded = atob(deepSearch(JSON.parse(data), 'stdout',"content").value);

          // console.log("img","/files"+deepSearch(JSON.parse(data), 'image/png',"relpath").value);

          
          setResult(deepSearch(JSON.parse(data), 'stdout',"content").value)
          if(+deepSearch(JSON.parse(data), 'image/png',"relpath") != null){
            setImgresult("/files"+deepSearch(JSON.parse(data), 'image/png',"relpath").value)
          }
          //console.log("imagepath",deepSearch(JSON.parse(data), 'image/png',"path"))
          // socket.disconnect(); // Disconnect from the server when receiving 'bye'
        });
        
        return () => socket.disconnect(); // Clean up on unmount
    }, []);



    const fetchPendingUpdates = async () => {
        try {
            const response = await fetch(window.location.origin +"/updates");
            const data = await response.json();
            console.log(data)
            setUpdates(data);
        } catch (error) {
            console.error('Error fetching pending updates:', error);
        }
    };


    const pingServer =async () =>{
    //   socket = io(window.location.origin);//'http://172.23.16.133:8080');
      socket.emit('message', "messageData");
      socket.emit("notice", "Starting");
    }
    const genImage =async () =>{
      //   socket = io(window.location.origin);//'http://172.23.16.133:8080');
        // socket.emit('message', "messageData");
        socket.emit("task", "genImage");
      }
      const runModule =async () =>{
        //   socket = io(window.location.origin);//'http://172.23.16.133:8080');
          // socket.emit('message', "messageData");
          //runsdxl sdxl:v0.9-lilypad1 PROMPT='moo'
          console.log( "task", selectedValue.split(" ")[0] + " -i "+ selectedValue.split(" ")[1] +`${promptValue}`)
          socket.emit("task", selectedValue.split(" ")[0] + " -i "+ selectedValue.split(" ")[1] +`${promptValue}`);
          //-i PromptEnv='PROMPT=big monkey'
        }
      
    return (
        <div>
            {/* <iframe src="http://localhost:3000/dashboard-solo/new?orgId=1&refresh=auto&from=1715712433852&to=1715714233852&panelId=1" width="450" height="200" frameborder="0"></iframe> */}
            <select value={selectedValue} onChange={handleSelectChange}>
                <option value="github.com/Lilypad-Tech/lilypad-module-sdxl:v0.9-lilypad1 PromptEnv=PROMPT=">github.com/Lilypad-Tech/lilypad-module-sdxl:v0.9-lilypad1</option>
                <option value="github.com/arsen3d/lilypad-module-lilysay:0.1.0 Message=">github.com/arsen3d/lilypad-module-lilysay:0.1.0</option>   
            </select>
            <input type="text" value={promptValue} onChange={handleChange } />
            <button onClick={runModule}>Run Module</button>
            {/* <br/>
            <button onClick={pingServer}>Cow Say</button>
            <button onClick={genImage}>SDXL Image Gen</button> */}
            <br/>
            Matches:
            <br/>
            <code>{matcher}</code>
            <br/>
            Deal Status:
            <br/>
            <code>{deal}</code>
            <br/>
            Result:
            <br/>
            <pre>{result}</pre>
            <br/>
            <br/>
            <img width={512} height={512} src={imgresult} />
            <br/>
            Messages:
            <br/>
            <code>{stdout}</code>
            <h1>Live Logs</h1>
            
            <ul>
                {updates.map((update) => (
                    <li key={update.id}> {new Date(update.timestamp).toLocaleString('en-US')}  {update.message}</li>
                ))}
            </ul>
        </div>
    );
};

export default App;
