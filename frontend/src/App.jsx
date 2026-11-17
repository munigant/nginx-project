import { useEffect, useState } from 'react'
import './App.css'

function App() {
 const [apiStatus, setApiStatus] = useState("Checking API...")

 useEffect(()=>{
  fetch("/api/health").then(response => {
    if(!response.ok){
      throw new Error('API request failed')
    }
    return response.json()
  })
  .then(data => {
    setApiStatus(`${data.service}: ${data.status}`)
  })
  .catch(()=>{
    setApiStatus("API unavailable")
  })
 },[])
 
  return (
    <div style={{padding:"40px", fontFamily:"sans-serif"}}>
      <h1>Devops Employee Dashboard</h1>
      <h2>Application Status</h2>
      <p>Go Backend: <strong>{apiStatus}</strong></p>
      <p>Frontend: React</p>
      <p>Web Server: Nginx</p>
    </div>
  )
}

export default App
