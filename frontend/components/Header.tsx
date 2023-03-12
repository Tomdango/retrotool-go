"use client";

import clsx from "clsx";
import React, { SyntheticEvent, useState } from "react";

const Header: React.FC = () => {
  const [isOpen, setIsOpen] = useState(false);

  const handleHeaderOpened = (event: SyntheticEvent) => {
    event.preventDefault();
    setIsOpen(!isOpen);
  };

  return (
    <div className="header header-fixed unselectable header-animated">
      <div className="header-brand">
        <div className="nav-item no-hover">
          <h6 className="title">Retrotool</h6>
        </div>
        <div
          className="nav-item nav-btn"
          id="header-btn"
          onClick={handleHeaderOpened}
        >
          <span></span> <span></span> <span></span>
        </div>
      </div>
      <div className={clsx("header-nav", { active: isOpen })} id="header-menu">
        <div className="nav-left">
          <div className="nav-item text-center">
            <a href="#">
              <span className="icon">
                <i className="fab fa-wrapper fa-twitter" aria-hidden="true"></i>{" "}
              </span>
            </a>
          </div>
        </div>
        <div className="nav-right">
          <div className="nav-item has-sub toggle-hover" id="dropdown">
            <a className="nav-dropdown-link">Menu</a>
            <ul className="dropdown-menu dropdown-animated" role="menu">
              <li role="menu-item">
                <a href="#">Profile</a>
              </li>
              <li role="menu-item">
                <a href="#">Log Out</a>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Header;
