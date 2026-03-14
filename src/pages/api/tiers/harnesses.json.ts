import type { APIRoute } from "astro";
import { getHarnessRankings } from "../../../data/compat";

export const GET: APIRoute = async () =>
  new Response(JSON.stringify(getHarnessRankings(), null, 2), {
    headers: { "content-type": "application/json; charset=utf-8" },
  });
