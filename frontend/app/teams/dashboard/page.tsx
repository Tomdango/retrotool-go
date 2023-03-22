import { NextPage } from "next";
import Link from "next/link";
import CreateTeamModal from "./components/CreateTeamModal";

const TeamsDashboardPage = async () => {
  return (
    <>
      <div className="content mt-10">
        <div className="level">
          <h1 className="title text-gray-700">Your Teams</h1>
          <a className="btn btn--lg" href="#create-team">
            New Team
          </a>
        </div>
        <div className="card">
          <div className="content px-5 py-2">
            <p>
              <b>Hmm, it looks a little empty in here.</b>
              <br />
              To get started, you&apos;ll need to{" "}
              <Link href="/teams/create">create a new team.</Link>
            </p>
          </div>
        </div>
      </div>
      <CreateTeamModal />
    </>
  );
};

export const getData = async () => {
  console.log("GETTING DATA");
};

export default TeamsDashboardPage;
