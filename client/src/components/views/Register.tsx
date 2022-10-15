import React, { useState } from 'react';
import { Button, Container, FormControl, Input, InputLabel } from '@mui/material';
import { styled } from '@mui/material/styles';
import { Link } from 'react-router-dom';
import { API } from '../../utils/constant';
import axios from 'axios';

const RegisterBtn = styled(Button)({
	color: '#000',
	border: '1px solid #000',
	fontWeight: 'bold',
	'&:hover': {
		color: '#fff',
		backgroundColor: '#000',
	}
});


function Register() {
  const [formInput, setFormInput] = useState({
    username: "",
    email: "",
    password: "",
    passwordCheck: ""
  })

	const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setFormInput((obj) => ({ ...obj, [e.target.name]: e.target.value ?? "" }));
	};

  const handleSubmit = async (e: React.SyntheticEvent) => {
    e.preventDefault()
    try {
      const res = await axios.post(`${API}/register`, formInput);
      console.log(res)
    } catch (err: any) {
      console.log(err.response)
    }
  }


	return (
		<Container maxWidth="xs" style={{
			display: 'flex',
			flexDirection: 'column',
			justifyContent: 'center',
			alignItems: 'center',
			minHeight: 'inherit'
    }}>
			<h1 style={{ margin: "2rem 0"}}>Register</h1>
      <form style={{
        display: 'flex',
				flexDirection: 'column',
				gap: '1rem',
				width: '100%'
      }}>
        <FormControl>
          <InputLabel htmlFor='username'>Username</InputLabel>
          <Input onChange={(e) => handleChange(e)} id="username" type="text" name="username" />
        </FormControl>
        <FormControl>
          <InputLabel htmlFor='email'>Email</InputLabel>
          <Input onChange={(e) => handleChange(e)} id="email" type="email" name="email" />
        </FormControl>
        <FormControl>
          <InputLabel htmlFor='password'>Password</InputLabel>
          <Input onChange={(e) => handleChange(e)} id="password" type="password" name="password" />
        </FormControl>
        <FormControl>
          <InputLabel htmlFor='passwordCheck'>Re-enter password</InputLabel>
          <Input onChange={(e) => handleChange(e)} id="passwordCheck" type="password" name="passwordCheck" />
        </FormControl>
        <RegisterBtn onClick={handleSubmit} type="submit">Register</RegisterBtn>
				<Link to="/">Sign in</Link>
      </form>
    </Container>
	);
}


export default Register;