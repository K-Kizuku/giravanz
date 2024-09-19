/* eslint-disable @typescript-eslint/ban-ts-comment */
import { FC, useState } from "react";

type Props = {
  ws: WebSocket;
};

type Message = {
  messages: {
    [key: string]: number;
  };
};

export const Recognition: FC<Props> = ({ ws }) => {
  const [word, setWord] = useState("test");
  const [result, setResult] = useState<Message>({} as Message);
  // @ts-ignore
  const SpeechRecognition = window.SpeechRecognition || webkitSpeechRecognition;
  const recognition = new SpeechRecognition();

  recognition.lang = "ja";
  recognition.continuous = true;
  // @ts-ignore
  recognition.onresult = (event) => {
    setWord(event.results[0][0].transcript);
    if (event.results[0][0].transcript === "") {
      recognition.stop();

      return;
    }

    ws.send(event.results[0][0].transcript);
    recognition.stop();
    // recognition.start();
  };
  recognition.onspeechend = () => {
    recognition.stop();
    console.log("end");
  };
  recognition.onspeechstart = () => {
    // recognition.start();
    console.log("start");
  };
  recognition.onend = () => {
    recognition.start();
    console.log("restart");
  };

  ws.onmessage = (event) => {
    console.log(event.data);
    const message = JSON.parse(event.data) as Message;
    if (!message.messages) {
      return;
    }
    setResult(message);
    // setResult((prev) => [event.data]);
  };

  // recognition.onspeechend = (event) => {
  //   recognition.stop();
  //   console.log("end");
  //   recognition.start();
  //   console.log("restart");
  // };
  console.log(result);
  return (
    <>
      <h1>{word}</h1>
      <button
        onClick={() => {
          recognition.start();
        }}
      >
        start
      </button>
      <div>
        {Object.keys(result.messages || {}).map((e) => {
          return (
            <div key={e} style={{ fontSize: `${14 + result.messages[e]}px` }}>
              {e}:{result.messages[e]}
            </div>
          );
        })}
      </div>
    </>
  );
};
