import React, { useState } from "react";
import Footer from "./Users/Layout/Footer";
import HeroSection from "./Users/Content/Homepage/components/HeroSection";
import InfoSection from "./Users/Content/Homepage/components/InfoSection";
import {
  homeObjOne,
  homeObjTwo,
  homeObjThree,
} from "./Users/Content/Homepage/components/InfoSection/Data";
import Navbar from "./Users/Layout/Navbar";
import Services from "./Users/Content/Homepage/components/Services";
import Sidebar from "./Users/Layout/Sidebar";

const Home = () => {
  const [isOpen, setIsOpen] = useState(false);

  const toggle = () => {
    setIsOpen(!isOpen);
  };

  return (
    <>
      <Sidebar isOpen={isOpen} toggle={toggle} />
      <Navbar toggle={toggle} />
      <HeroSection />
      <InfoSection {...homeObjOne} />
      <InfoSection {...homeObjTwo} />
      <Services />
      <InfoSection {...homeObjThree} />
      <Footer />
    </>
  );
};

export default Home;
