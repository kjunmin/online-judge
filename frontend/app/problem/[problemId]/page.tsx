
type ProblemProps = {
  params: {
    problemId: string
  }
}

export default function Problem({params}: ProblemProps) {
  return (
    <div className="flex flex-col items-center gap-3">
      <section>
        Problem page: {params.problemId}
      </section>
    </div>
  );
}