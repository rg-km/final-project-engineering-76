import { useState, useEffect } from 'react';
import '/node_modules/bootstrap/dist/css/bootstrap.min.css';
import React from 'react';
import { Link } from 'react-router-dom'; 

function Register() 
{
    const [name, setName] = useState('');
    const [password, setPassword] = useState('');
    const [role, setRole] = useState('');
    const options = [
        {value: "user", label:"User"},
        {value: "admin", label:"Admin"}
    ];
    const [redirect, setRedirect] = useState(false);
    const [submitted, setSubmitted] = useState(false);
    const [error, setError] = useState(false);

    const handleName = (e) => {
        setName(e.target.value);
        setSubmitted(false);
    };
    const handleRole = (e) => {
        setRole(e.target.value);
        setSubmitted(false);
    };
    const handlePassword = (e) => {
        setPassword(e.target.value);
        setSubmitted(false);
    };
    const handleSubmit = (e) => {
        e.preventDefault();

        if(name === '' || password === '' || role === '') {
            alert('Please fill in all the fields');
            setError(true);
        } else {
            let ob ={
                name: name,
                password: password,
                role: role,
            }
    
            let oldData = localStorage.getItem('data');
            if (oldData==null){
                oldData = [];
                oldData.push(ob);
                localStorage.setItem('data', JSON.stringify(oldData));
            } else{
                let oldArr = JSON.parse(oldData);
                oldArr.push(ob);
                localStorage.setItem('data', JSON.stringify(oldArr));
                console.log(oldArr,'hhg');
            }
    
            setName('');
            setPassword('');
            setRole('');
            console.log({name, password, role});
            setRedirect(true);
        }

        

    };

    const errorMessage = () => {
        if (error) {
            return <p>Please fill in fields</p>;
        }
    };

    const submit = async (e) => {
        e.preventDefault();

        const response = await fetch(`http://localhost:2213/api/signup`, {
            method: 'POST',
            mode:'no-cors',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name,
                password,
                role
            })
        });

        const content = await response.json();
        console.log(content);   
        
        setRedirect(true);
    }

    if (redirect) {
        return <Link to="/login" />;
    }

    return (
        <div className='register d-flex row justify-content-center py-5'>
            
                <div className='register-form col-md-6 p-4'>
                    <h1 className='mb-3'>Register</h1>
                    <div className='error-nofill'>
                        <p onChange={errorMessage}></p>
                    </div>
                    <form>
                        <div className='form-group mb-3'>
                            <label>Nama</label>
                            <input
                                type='text'
                                className='form-control'
                                value={name}
                                placeholder='Masukkan Nama'
                                onChange={handleName}
                            />
                        </div>
                        
                            <div className='form-group mb-3'>
                                <label>Role</label>
                                <select className='form-control' onClick={handleRole}>
                                    {options.map(option => (
                                        <option key={option.value} value={option.value}>
                                            {option.label}
                                        </option>
                                    ))}
                                </select>
                            </div>
                        
                        <div className='form-group mb-3'>
                            <label>Password</label>
                            <input
                                type='password'
                                className='form-control'
                                value={password}
                                placeholder='Masukkan Password'
                                onChange={handlePassword}
                            />
                        </div>
                        <div className='row mt-5'>
                            <div className='col'>
                                <button type='button' className='btn btn-register p-1'>Cancel
                                        <Link to='/login' className='text-white'></Link>
                                </button>
                            </div>
                            <div className='col'>
                                <button type='submit' className='btn btn-register p-1' onClick={handleSubmit}>Submit</button>
                            </div>
                        </div>
                    </form>
                </div>
            
        </div>
        
    )
}

export default Register;