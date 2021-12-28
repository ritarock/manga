import Link from 'next/link'

export default function ReleaseList() {
  const dt = new Date()
  const yyyymmList: string[] = []

  for (let i=0;i<6;i++){
    yyyymmList.push(formatDate(
      subMonth(
        dt.getFullYear(),
        +("00"+(dt.getMonth()+1)).slice(-2),
        i
      )
    ))
  }
  return (
    <>
      <div>
        {yyyymmList.map(dt => (
          <div key={dt}>
            <Link
              href={{
                pathname: "/manga/[date]",
                query: {date: `${dt}`}
              }}
            >
              <a>{dt.slice(0,4)}/{dt.slice(-2)}</a>
            </Link>
          </div>
        ))}
      </div>
    </>
  )
}

function subMonth(year: number, month: number, sub:number) {
  return new Date(year, month-sub)
}

function formatDate(date: Date) {
  return `${date.getUTCFullYear()}${("00"+(date.getUTCMonth()+1)).slice(-2)}`
}
