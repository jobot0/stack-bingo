import "./App.css";

import React, { useState } from "react";

import logo from "./logo.svg";
import { SaveDumbUsecase } from "../core/SaveDumbUsecase";
import { RestDumbRespository } from "../data/RestDumbRepository";

const App: React.FunctionComponent = () => {
  const restDumbRepository = new RestDumbRespository();
  const saveDumbUsecase = new SaveDumbUsecase(restDumbRepository);
  const onDumbClick = async (el: any) => {
    el.preventDefault();
    await saveDumbUsecase.execute("toto");
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p onClick={onDumbClick}>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <button onClick={onDumbClick}>Toto</button>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
};

export default App;
