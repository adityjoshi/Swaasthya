import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import ReCAPTCHA from "react-google-recaptcha";
import NavigationLog from '../components/NavigationLog';
import './Login.css'; 

const Login = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();
    
    // Simple validation
    if (!email || !password) {
      setError('Please fill in all fields');
      return;
    }

    // Add your login logic here
    try {
      // Example login request
      // const response = await loginUser({ email, password });
      // if (response.success) {
      navigate('/otp-login'); // Or any other route
      // } else {
      //   setError(response.message);
      // }
    } catch (err) {
      setError('Login failed. Please try again.');
    }
  };

  return (
    <>
      <NavigationLog />
      <div className="login-container">
        <form onSubmit={handleSubmit} className="login-form">
          <img src="/assets/SIH_logo2.png" alt="Logo" className="form-logo" />
          <h3>Log In</h3>
          
          {error && <div className="alert alert-danger">{error}</div>}

          <div className="mb-3">
            <label>Email address</label>
            <input
              type="email"
              className="form-control"
              placeholder="Enter email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>

          <div className="mb-3">
            <label>Password</label>
            <input
              type="password"
              className="form-control"
              placeholder="Enter password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>

          <div className="recaptcha-container">
            <ReCAPTCHA
              sitekey="6LdR1jMqAAAAAHVJopbBevQmLVJdsNik_LcjBPXG" // Replace with your ReCAPTCHA site key
              onChange={(value) => console.log("Captcha value:", value)}
            />
          </div>

          <div className="mb-3">
            <div className="custom-control custom-checkbox">
              <input
                type="checkbox"
                className="custom-control-input"
                id="customCheck1"
              />
              <label className="custom-control-label" htmlFor="customCheck1">
                Remember me
              </label>
            </div>
          </div>

          <div className="d-grid">
            <button type="submit" className="submit-button">
              Submit
            </button>
          </div>
          <p className="forgot-password text-right">
            Forgot <a href="#">password?</a>
          </p>
        </form>
      </div>
    </>
  );
};

export default Login;
