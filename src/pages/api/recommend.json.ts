import type { APIRoute } from "astro";
import { getRecommendedCombos } from "../../data/compat";

export const GET: APIRoute = async ({ url }) => {
  const usecase = url.searchParams.get("usecase") ?? undefined;
  const recommendations = getRecommendedCombos(usecase, 5);

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
