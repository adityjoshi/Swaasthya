// import React, { useState } from 'react';
// import { useNavigate } from 'react-router-dom';
// import NavigationLog from '../components/NavigationLog';
// // import './signUp.css';

// const Hospital_SignUp = () => {
//   const [department, setDepartment] = useState("");
//   const [otherDepartment, setOtherDepartment] = useState("");
//   const navigate = useNavigate();

//   const handleDepartmentChange = (e) => {
//     setDepartment(e.target.value);
//     if (e.target.value === "other") {
//       setOtherDepartment(""); 
//     }
//   };

//   const handleSubmit = (event) => {
//     event.preventDefault();
//     // Add sign-up logic here

//     // On successful sign-up, navigate to another page
//     navigate('/hospital-login'); // Or any other route
//   };

//   return (
//     <>
//       <NavigationLog />
//       <div className="signup-container">
//         <form className="signup-form" onSubmit={handleSubmit}>
//           <img src="/assets/SIH_logo2.png" alt="Logo" className="form-logo" />
//           <h3>Sign Up</h3>
//           <div className="mb-3">
//             <label>Name</label>
//             <input type="text" className="form-control" placeholder="Enter your Name" required />
//           </div>

//           <div className="mb-3">
//             <label>Hospital ID</label>
//             <input type="number" className="form-control" placeholder="Enter your Hospital ID" required />
//           </div>

//           <div className="mb-3">
//             <label>Address</label>
//             <input type="text" className="form-control" placeholder="Enter your Address" required />
//           </div>

//           <div className="mb-3">
//             <label>Email</label>
//             <input type="email" className="form-control" placeholder="Enter email" required />
//           </div>

//           <div className="mb-3">
//             <label>Website</label>
//             <input type="text" className="form-control" placeholder="Website (If any)" />
//           </div>

//           <div className="mb-3">
//             <label>Department</label>
//             <select 
//               className="form-control" 
//               id="department" 
//               value={department}
//               onChange={handleDepartmentChange}
//               required
//             >
//               <option value="">Choose...</option>
//               <option value="Orthopaedic">Orthopaedic</option>
//               <option value="Cardiac">Cardiac</option>
//               <option value="Physician">Physician</option>
//               <option value="Pediatrician">Pediatrician</option>
//               <option value="ENT">ENT</option>
//               <option value="Dermatologist">Dermatologist</option>
//               <option value="Gynecologist">Gynecologist</option>
//               <option value="Dentist">Dentist</option>
//               <option value="other">Other</option>
//             </select>
//             {department === "other" && (
//               <div className="mt-3">
//                 <label>Specify Department</label>
//                 <input
//                   type="text"
//                   className="form-control"
//                   placeholder="Specify your department"
//                   value={otherDepartment}
//                   onChange={(e) => setOtherDepartment(e.target.value)}
//                   required
//                 />
//               </div>
//             )}
//           </div>

//           <div className="d-grid">
//             <button type="submit" className="submit-button">
//               Sign Up
//             </button>
//           </div>

//           <p className="forgot-password text-right">
//             Already registered <a href="/hospital-login">sign in?</a>
//           </p>
//         </form>
//       </div>
//     </>
//   );
// };

// export default Hospital_SignUp;


import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios'; // Import axios
import NavigationLog from '../components/NavigationLog';
import './Signup.css'; // Ensure you have the correct CSS file if needed

// const Hospital_SignUp = () => {
//   // State variables
//   const [name, setName] = useState('');
//   const [hospitalID, setHospitalID] = useState('');
//   const [address, setAddress] = useState('');
//   const [email, setEmail] = useState('');
//   const [website, setWebsite] = useState('');
//   const [department, setDepartment] = useState('');
//   const [otherDepartment, setOtherDepartment] = useState('');
//   const [error, setError] = useState('');
//   const navigate = useNavigate();

//   const handleDepartmentChange = (e) => {
//     setDepartment(e.target.value);
//     if (e.target.value === "other") {
//       setOtherDepartment(""); 
//     }
//   };

//   const handleSubmit = async (event) => {
//     event.preventDefault();

//     try {
//       const response = await axios.post(`http://localhost:2426/hospitaladmin`, {
//         Name: name,
//         HospitalID: hospitalID,
//         Address: address,
//         Email: email,
//         Website: website,
//         Department: department === "other" ? otherDepartment : department
//       });

//       if (response.status === 200) {
//         navigate('/hospital-login'); // Redirect on success
//       } else {
//         setError('Registration failed. Please try again.');
//       }
//     } catch (err) {
//       setError('Registration failed. Please try again.');
//     }
//   };

//   return (
//     <>
//       <NavigationLog />
//       <div className="signup-container">
//         <form className="signup-form" onSubmit={handleSubmit}>
//           <img src="/assets/SIH_logo2.png" alt="Logo" className="form-logo" />
//           <h3>Sign Up</h3>
          
//           {error && <div className="alert alert-danger">{error}</div>}

//           <div className="mb-3">
//             <label>Name</label>
//             <input 
//               type="text" 
//               className="form-control" 
//               placeholder="Enter your Name" 
//               value={name}
//               onChange={(e) => setName(e.target.value)}
//               required 
//             />
//           </div>

//           <div className="mb-3">
//             <label>Hospital ID</label>
//             <input 
//               type="number" 
//               className="form-control" 
//               placeholder="Enter your Hospital ID" 
//               value={hospitalID}
//               onChange={(e) => setHospitalID(e.target.value)}
//               required 
//             />
//           </div>

//           <div className="mb-3">
//             <label>Address</label>
//             <input 
//               type="text" 
//               className="form-control" 
//               placeholder="Enter your Address" 
//               value={address}
//               onChange={(e) => setAddress(e.target.value)}
//               required 
//             />
//           </div>

//           <div className="mb-3">
//             <label>Email</label>
//             <input 
//               type="email" 
//               className="form-control" 
//               placeholder="Enter email" 
//               value={email}
//               onChange={(e) => setEmail(e.target.value)}
//               required 
//             />
//           </div>

//           <div className="mb-3">
//             <label>Website</label>
//             <input 
//               type="text" 
//               className="form-control" 
//               placeholder="Website (If any)" 
//               value={website}
//               onChange={(e) => setWebsite(e.target.value)}
//             />
//           </div>

//           <div className="mb-3">
//             <label>Department</label>
//             <select 
//               className="form-control" 
//               id="department" 
//               value={department}
//               onChange={handleDepartmentChange}
//               required
//             >
//               <option value="">Choose...</option>
//               <option value="Orthopaedic">Orthopaedic</option>
//               <option value="Cardiac">Cardiac</option>
//               <option value="Physician">Physician</option>
//               <option value="Pediatrician">Pediatrician</option>
//               <option value="ENT">ENT</option>
//               <option value="Dermatologist">Dermatologist</option>
//               <option value="Gynecologist">Gynecologist</option>
//               <option value="Dentist">Dentist</option>
//               <option value="other">Other</option>
//             </select>
//             {department === "other" && (
//               <div className="mt-3">
//                 <label>Specify Department</label>
//                 <input
//                   type="text"
//                   className="form-control"
//                   placeholder="Specify your department"
//                   value={otherDepartment}
//                   onChange={(e) => setOtherDepartment(e.target.value)}
//                   required
//                 />
//               </div>
//             )}
//           </div>

//           <div className="d-grid">
//             <button type="submit" className="submit-button">
//               Sign Up
//             </button>
//           </div>

//           <p className="forgot-password text-right">
//             Already registered <a href="/hospital-login">sign in?</a>
//           </p>
//         </form>
//       </div>
//     </>
//   );
// };

// export default Hospital_SignUp;

const Hospital_SignUp = () => {
  // State variables
  const [fullName, setFullName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [contactNumber, setContactNumber] = useState('');
  const [userType, setUserType] = useState('admin'); // Default to 'admin'
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await axios.post('http://localhost:2426/hospitaladmin', {
        full_name: fullName,
        email,
        password,
        contact_number: contactNumber,
        user_type: userType,
      });

      if (response.status === 201) {
        navigate('/hospital-login'); // Redirect on success
      } else {
        setError('Registration failed. Please try again.');
      }
    } catch (err) {
      setError('Registration failed. Please try again.');
    }
  };

  return (
    <>
      <NavigationLog />
      <div className="signup-container">
        <form className="signup-form" onSubmit={handleSubmit}>
          <img src="/assets/SIH_logo2.png" alt="Logo" className="form-logo" />
          <h3>Sign Up</h3>
          
          {error && <div className="alert alert-danger">{error}</div>}

          <div className="mb-3">
            <label>Full Name</label>
            <input 
              type="text" 
              className="form-control" 
              placeholder="Enter your Full Name" 
              value={fullName}
              onChange={(e) => setFullName(e.target.value)}
              required 
            />
          </div>

          <div className="mb-3">
            <label>Email</label>
            <input 
              type="email" 
              className="form-control" 
              placeholder="Enter your Email" 
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
              placeholder="Enter your Password" 
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required 
            />
          </div>

          <div className="mb-3">
            <label>Contact Number</label>
            <input 
              type="text" 
              className="form-control" 
              placeholder="Enter your Contact Number" 
              value={contactNumber}
              onChange={(e) => setContactNumber(e.target.value)}
              required 
            />
          </div>

          <div className="mb-3">
            <label>User Type</label>
            <select 
              className="form-control" 
              value={userType}
              onChange={(e) => setUserType(e.target.value)}
              required
            >
              <option value="admin">Admin</option>
              <option value="staff">Staff</option>
            </select>
          </div>

          <div className="d-grid">
            <button type="submit" className="submit-button">
              Sign Up
            </button>
          </div>

          <p className="forgot-password text-right">
            Already registered <a href="/hospital-login">sign in?</a>
          </p>
        </form>
      </div>
    </>
  );
};

export default Hospital_SignUp;
