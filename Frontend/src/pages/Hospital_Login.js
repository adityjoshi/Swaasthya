// import React, { useState } from 'react';
// import { useNavigate } from 'react-router-dom';
// import ReCAPTCHA from "react-google-recaptcha";
// import NavigationLog from '../components/NavigationLog';
// import './Login.css';

// const Hospital_Login = () => {
//   const navigate = useNavigate();
//   const [number, setNumber] = useState('');
//   const [password, setPassword] = useState('');

//   const handleSubmit = (event) => {
//     event.preventDefault();
//     // Add login logic here

//     // On successful login, navigate to another page
//     navigate('/otp-login'); // Or any other route
//   };

//   return (
//     <>
//       <NavigationLog />
//       <div className="login-container">
//         <form onSubmit={handleSubmit} className="login-form">
//           <img src="../assets/SIH_logo2.png" alt="Logo" className="form-logo" />
//           <h3>Log In</h3>

//           <div className="mb-3">
//             <label>Hospital ID</label>
//             <input
//               type="number"
//               className="form-control"
//               placeholder="Enter your ID"
//               value={number}
//               onChange={(e) => setNumber(e.target.value)}
//             />
//           </div>

//           <div className="mb-3">
//             <label>Password</label>
//             <input
//               type="password"
//               className="form-control"
//               placeholder="Enter password"
//               value={password}
//               onChange={(e) => setPassword(e.target.value)}
//             />
//           </div>

//           <div className="recaptcha-container">
//             <ReCAPTCHA
//               sitekey="6LdR1jMqAAAAAHVJopbBevQmLVJdsNik_LcjBPXG" // Replace with your ReCAPTCHA site key
//               onChange={(value) => console.log("Captcha value:", value)}
//             />
//           </div>

//           <div className="mb-3">
//             <div className="custom-control custom-checkbox">
//               <input
//                 type="checkbox"
//                 className="custom-control-input"
//                 id="customCheck1"
//               />
//               <label className="custom-control-label" htmlFor="customCheck1">
//                 Remember me
//               </label>
//             </div>
//           </div>

//           <div className="d-grid">
//             <button type="submit" className="submit-button">
//               Submit
//             </button>
//           </div>
//           <p className="forgot-password text-right">
//             Forgot <a href="#">password?</a>
//           </p>
//         </form>
//       </div>
//     </>
//   );
// };

// export default Hospital_Login;


import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import ReCAPTCHA from 'react-google-recaptcha';
import axios from 'axios'; 
import NavigationLog from '../components/NavigationLog';
import './Login.css';

const Hospital_Login = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();

    // Add login logic here. Example:
    try {
      // Assuming you send a POST request for login
      const response = await axios.post('http://localhost:2426/adminLogin', {
        email,
        password
      });

      if (response.status === 200) {
        navigate('/otp-login'); // Or any other route upon successful login
      } else {
        // Handle errors or unsuccessful login
        console.error('Login failed. Please try again.');
      }
    } catch (err) {
      console.error('Login failed. Please try again.', err);
    }
  };

  return (
    <>
      <NavigationLog />
      <div className="login-container">
        <form onSubmit={handleSubmit} className="login-form">
          <img src="../assets/SIH_logo2.png" alt="Logo" className="form-logo" />
          <h3>Log In</h3>

          <div className="mb-3">
            <label>Email</label>
            <input
              type="email"
              className="form-control"
              placeholder="Enter your email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
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
              required
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

export default Hospital_Login;
