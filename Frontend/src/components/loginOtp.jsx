import React, { useState } from 'react';
import OTPInput from 'react-otp-input';
import { ColorRing } from 'react-loader-spinner';
import './loginOtp.css';

const OtpLogin = ({ onLogin }) => {
    const [email, setEmail] = useState('');
    const [otp, setOtp] = useState(['', '', '', '']);
    const [error, setError] = useState('');
    const [isEmailVerified, setIsEmailVerified] = useState(false);
    const [loading, setLoading] = useState(false);

    const validateEmail = (email) => {
        return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
    };

    const handleEmailVerification = () => {
        if (validateEmail(email)) {
            setLoading(true);
            setTimeout(() => {
                setIsEmailVerified(true);
                setLoading(false);
                setError('');
            }, 2000);
        } else {
            setError('Invalid Email');
        }
    };

    const handleOTPVerification = () => {
        const enteredOTP = otp.join('');
        if (/^\d{4}$/.test(enteredOTP)) {
            const username = email.split('@')[0].charAt(0).toUpperCase() + email.split('@')[0].slice(1);
            onLogin(username);
        } else {
            setError('Incorrect OTP');
        }
    };

    return (
        <div className='otp-container'>
            {!isEmailVerified ? (
                <>
                    <h1 className='otp-heading'>Enter Your Email:</h1>
                    <input
                        type='email'
                        placeholder='test@example.com'
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        className='email-input'
                        aria-label="Email"
                    />
                    {error && <span className='otp-error'>{error}</span>}
                    <button 
                        onClick={handleEmailVerification} 
                        className='otp-verify-button'
                        aria-live="polite"
                    >
                        {loading ? (
                            <ColorRing color="#fff" height={50} width={50} />
                        ) : (
                            'Verify Email'
                        )}
                    </button>
                </>
            ) : (
                <>
                    <h1 className='otp-heading'>Enter Your OTP:</h1>
                    <div className='otp-input-container'>
                        <OTPInput
                            value={otp.join('')}
                            onChange={(value) => setOtp(value.split(''))}
                            numInputs={4}
                            isInputNum
                            inputStyle={{
                                width: '4rem',
                                height: '4rem',
                                margin: '0 0.5rem',
                                fontSize: '2rem',
                                borderRadius: '8px',
                                border: '2px solid #ced4da'
                            }}
                            renderInput={(inputProps, index) => <input {...inputProps} key={index} />}
                        />
                    </div>
                    {error && <span className='otp-error'>{error}</span>}
                    <button 
                        onClick={handleOTPVerification} 
                        className='otp-login-button'
                        aria-live="polite"
                    >
                        Login
                    </button>
                </>
            )}
        </div>
    );
};

const Challenge41 = () => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [username, setUsername] = useState('');

    const handleLogin = (username) => {
        setIsLoggedIn(true);
        setUsername(username);
    };

    const handleLogout = () => {
        setIsLoggedIn(false);
        setUsername('');
    };

    return (
        <section className='landing-section'>
            <div className='landing-div'>
                {isLoggedIn ? (
                    <div className='welcome-container'>
                        <h1 className='welcome-heading'>Welcome, {username}</h1>
                        <button onClick={handleLogout} className='logout-button'>
                            Logout
                        </button>
                    </div>
                ) : (
                    <OtpLogin onLogin={handleLogin} />
                )}
            </div>
        </section>
    );
};

export default Challenge41;
