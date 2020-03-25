import React, { useState } from "react";
import { HelloWorldClient } from "./proto/helloworld/helloworldServiceClientPb";
import {
  SayHelloRequest,
  CurrentTimeRequest,
} from "./proto/helloworld/helloworld_pb";
import grpcWeb from "grpc-web";
import {
  Navbar,
  Container,
  Row,
  Col,
  Label,
  Input,
  Button,
  FormGroup,
} from "reactstrap";

let target = "http://localhost:8888";
if (process.env.REACT_APP_HELLOWORLD_API) {
  target = process.env.REACT_APP_HELLOWORLD_API;
}
let client = new HelloWorldClient(target);

function useInputState(
  initialState: string
): [string, (e: React.ChangeEvent<HTMLInputElement>) => void] {
  const [value, setValue] = useState(initialState);
  function onChange(e: React.ChangeEvent<HTMLInputElement>) {
    setValue(e.target.value);
  }
  return [value, onChange];
}

function App() {
  const [language, onLanguageChange] = useInputState("");
  const [languageResponse, setLanguageResponse] = useState("");

  function sayHello() {
    let req = new SayHelloRequest();
    req.setLanguage(language);
    let stream = client.sayHello(req, null, () => {});
    stream.on("data", (resp) => {
      setLanguageResponse(resp.getHello());
    });
    stream.on("error", (err: grpcWeb.Error) => {
      console.error(err);
      setLanguageResponse("error: " + err.message);
    });
  }

  const [currentTimeResponse, setCurrentTimeResponse] = useState("");

  function currentTime() {
    let req = new CurrentTimeRequest();
    let stream = client.currentTime(req, null, () => {});
    stream.on("data", (resp) => {
      let t = resp.getCurrentTime();
      if (t) {
        let d = t.toDate();
        setCurrentTimeResponse(d.toString());
      } else {
        setCurrentTimeResponse("undefined");
      }
    });
    stream.on("error", (err: grpcWeb.Error) => {
      console.error(err);
      setCurrentTimeResponse("error: " + err.message);
    });
  }

  return (
    <Container>
      <Navbar>Protip gRPC-Web Example</Navbar>
      <Row>
        <Col>
          <p>
            gRPC-Web target is currently <code>{target}</code>.
          </p>
          <p>
            Ensure the gRPC-Web server is running. From the <code>golang</code>{" "}
            directory run{" "}
            <code>go run github.com/protip-dev/examples/cmd/grpcwebserver</code>
            .
          </p>
          <p>
            You can change the gRPC-Web target via the{" "}
            <code>REACT_APP_HELLOWORLD_API</code> environment variable when
            running <code>yarn start</code>.
          </p>
        </Col>
      </Row>
      <Row>
        <Col>
          <h4>SayHello()</h4>
          <FormGroup>
            <Label for="language">Language</Label>
            <Input
              type="text"
              value={language}
              onChange={onLanguageChange}
              placeholder={
                'Two-letter ISO 639-1 language code, defaults to "en"'
              }
            />
          </FormGroup>
          <FormGroup>
            <Button onClick={sayHello}>Send RPC</Button>
          </FormGroup>

          <FormGroup>
            <Input type="textarea" value={languageResponse} />
          </FormGroup>
        </Col>
      </Row>

      <Row>
        <Col>
          <h4>CurrentTime()</h4>
          <FormGroup>
            <Button onClick={currentTime}>Send RPC</Button>
          </FormGroup>

          <FormGroup>
            <Input type="textarea" value={currentTimeResponse} />
          </FormGroup>
        </Col>
      </Row>
    </Container>
  );
}

export default App;
