import React, { useState } from 'react';
import LoginForm from './components/LoginForm';
import './index.css';

function App() {
  const user = {
    name: 'John',
    password: '123456'
  }

  const [login, setLogin] = useState({ name: '', password: '' })
  const [error, setError] = useState('')

  const Login = details => {
    console.log(details)

    if (details.name === user.name && details.password === user.password) {
      console.log('Login success')
      setLogin({
        name: details.name,
        password: details.password
      })
    } else {
      console.log('Login failed')
      setError('Login failed!')
    }
  } // end of Login
  
  const Logout = () => {
    setLogin({
      name: '',
      password: ''
    })
  } // end of Logout

  return (
    <div className="App">
      {(login.name != "") ? (
        <div className="welcome">
          <h2>Welcome, <span>{login.name}</span></h2>
          <button className='logout' onClick={Logout}>Logout</button>
          </div>
      ) : (
          <LoginForm Login={Login} error={error} />
          )}
      </div>
  );
}

export default App;
