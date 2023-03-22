"use client";

import Button from "@/components/cirrus/Button";
import Input from "@/components/cirrus/Input";
import Modal from "@/components/cirrus/Modal";
import { useAuth } from "@/hooks/UseAuth";
import { FormEvent, useState } from "react";

const CreateTeamModal = () => {
  const [teamName, setTeamName] = useState("");
  const auth = useAuth();

  const handleSubmit = async (event: FormEvent) => {
    event.preventDefault();
    if (!auth.isLoggedIn) {
      return;
    }

    console.log(auth.session);
    const response = await fetch(
      "https://api.retrotool.tomjc.dev/teams/create",
      {
        method: "POST",
        headers: new Headers({
          Authorization: `Bearer ${auth.session.accessToken}`,
        }),
      }
    );

    console.log(response.statusText);
    console.log(await response.text());
  };

  return (
    <Modal id="create-team" animation="zoom-in">
      <Modal.Overlay />
      <Modal.Content className="min-w-xs">
        <Modal.Header>
          <Modal.CloseIcon />
          <Modal.Title>New Team</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <p></p>
          <form onSubmit={handleSubmit}>
            <Input
              placeholder="Team Name"
              value={teamName}
              onChange={(e) => setTeamName(e.currentTarget.value)}
            />
            <Button type="submit">Create</Button>
          </form>
        </Modal.Body>
      </Modal.Content>
    </Modal>
  );
};

export default CreateTeamModal;
