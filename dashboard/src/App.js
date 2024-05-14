import React, { useEffect, useState } from 'react';
import io from 'socket.io-client';
const ENDPOINT = 'http://localhost:8000'; // Change this to your server's endpoint

const App = () => {
    const [updates, setUpdates] = useState([]);
    const [response, setResponse] = useState('');
    const [stdout, setStdout] = useState('');
    const [matcher, setMatcher] = useState('');
    const [deal, setDeal] = useState('');
    const [result, setResult] = useState('');
    
    const socket = io({transports: ['websocket']});
    // let socket;
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
          setResult(data)
          //setResponse(data);
        });
        socket.on('reply', data => {
          console.log('Received reply:', data);
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
      socket.emit("notice", "world");
    }
    return (
        <div>
            <button onClick={pingServer}>Cow Say</button>
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
