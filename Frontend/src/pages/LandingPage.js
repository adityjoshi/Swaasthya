import React, { useState, useEffect } from 'react';
import CommonNavigation from '../components/CommonNavigation';
import './LandingPage.css';

const LandingPage = () => {
    return (
        <div>
            <CommonNavigation />
            <div className="welcome-section">
                <h1>Welcome to <span className="highlight-text">Swaasthya</span></h1>
                <p>
                    "Transforming <span className="highlight-text">Healthcare</span> with Seamless Management and Enhanced Patient Care."
                </p>
            </div>

            <div className="full-cover-photo"></div>

            <section className="about-us">
                <div className="container">
                    <h2>About Us</h2>
                    <p>
                        At <span className="highlight-text">Swaasthya</span>, we are committed to revolutionizing healthcare management through innovative technology. Our platform seamlessly integrates patient care, OPD management, and inventory control to ensure efficient and high-quality healthcare services. With a focus on enhancing patient experience and optimizing hospital operations, Swaasthya leverages advanced queuing models and real-time data to streamline processes and improve outcomes. Our mission is to provide hospitals and healthcare providers with the tools they need to deliver exceptional care and manage resources effectively. Join us in transforming healthcare for a healthier tomorrow.
                    </p>
                </div>
            </section>


            <section className="outro-line"></section>

            <footer className="footer">
                <div className="footer-content">
                    <div className="footer-logo">
                        <a href="/">
                            <img src="/assets/SIH_logo2.png" alt="Logo" width="40" height="32" />
                        </a>
                    </div>
                    <ul className="footer-nav">
                        <li><a href="#" className="footer-link">Privacy Policy</a></li>
                        <li><a href="#" className="footer-link">Terms of Service</a></li>
                        <li><a href="#" className="footer-link">Contact Us</a></li>
                    </ul>
                    <div className="footer-social">
                        <a href="#" className="social-link">
                            <img src="assets/facebook.png" alt="Facebook" width="24" height="24" />
                        </a>
                        <a href="#" className="social-link">
                            <img src="assets/instagram.png" alt="Twitter" width="24" height="24" />
                        </a>
                        <a href="#" className="social-link">
                            <img src="assets/twitter.png" alt="Instagram" width="24" height="24" />
                        </a>
                    </div>
                </div>
                <div className="footer-bottom">
                    <p>&copy; 2024 Swaasthya. All rights reserved.</p>
                </div>
            </footer>
        </div>
    );
};

export default LandingPage;
