import React, { useState } from "react";

function LoginForm({ Login, error }) {
    const [details, setDetails] = useState({ name: '', password: '' })
    
    const submit = e => {
        e.preventDefault();
            Login(details);
        
    }
    return (
            <form onSubmit= {submit}>
                <div className="form-inner">
                    <h2>Login</h2>
                    {(error != "") ? (<div className="error">{error}</div>) : (<div></div>)}
                    <div className="form-group">
                        <label htmlFor="name">Name:</label>
                    <input type="text" name="name" id="name" onChange={e => setDetails({ ...details, name: e.target.value })} value={details.name}
                    />
                    </div>
                    <div className="form-group">
                        <label htmlFor="password">Password:</label>
                        <input type="password" name="password" id="password" onChange={e => setDetails({ ...details, password: e.target.value })} value={details.password}/>
                </div>
                <input className="login-btn" type="submit" value="LOGIN" />
                <div>
                    <br />
                    <input className="regis-btn" type="submit" value="REGISTRASI" /></div>
            
                </div>
            </form>
    )
}

export default LoginForm;