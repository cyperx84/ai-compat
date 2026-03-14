import type { APIRoute } from "astro";
import { getModelRankings } from "../../../data/compat";

export const GET: APIRoute = async () =>
  new Response(JSON.stringify(getModelRankings(), null, 2), {
    headers: { "content-type": "application/json; charset=utf-8" },
  });
