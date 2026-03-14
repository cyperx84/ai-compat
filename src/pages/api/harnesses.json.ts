import type { APIRoute } from "astro";
import { compatData } from "../../data/compat";

export const GET: APIRoute = async () =>
  new Response(JSON.stringify(compatData.harnesses, null, 2), {
    headers: { "content-type": "application/json; charset=utf-8" },
  });
