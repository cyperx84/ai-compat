import type { APIRoute } from "astro";
import { getRecommendedCombos } from "../../data/compat";

export const GET: APIRoute = async () =>
  new Response(JSON.stringify(getRecommendedCombos(undefined, 100), null, 2), {
    headers: { "content-type": "application/json; charset=utf-8" },
  });
