import React from 'react';
import './App.css';
import Login from './components/views/Login';
import Register from './components/views/Register';
import { BrowserRouter, Routes, Route } from 'react-router-dom';

function App() {
	return (
		<div className="App">
			<main style={{
				minHeight: 'inherit'
			}}>
				<BrowserRouter>
					<Routes>
						<Route path="/">
							<Route index element={<Login />} />
							<Route path="register" element={<Register />} />
						</Route>
					</Routes>
				</BrowserRouter>
			</main>
		</div>
	);
}

export default App;
