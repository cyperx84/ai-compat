import type { APIRoute } from "astro";
import { getBestCombos } from "../../data/compat";

export const GET: APIRoute = async ({ url }) => {
  const usecase = url.searchParams.get("usecase") ?? undefined;
  const limit = Number(url.searchParams.get("limit") ?? "5");
  const recommendations = getBestCombos(usecase, Number.isNaN(limit) ? 5 : limit);

  return new Response(
    JSON.stringify(
      {
        usecase: usecase ?? "all",
        count: recommendations.length,
        recommendations,
      },
      null,
      2,
    ),
    {
      headers: { "content-type": "application/json; charset=utf-8" },
    },
  );
};
