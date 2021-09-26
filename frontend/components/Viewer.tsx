type Book = {
	isbn: string
	title: string
	publisher: string
	pubdate: string
	cover: string
	author: string
	subjectCode: string
}

export default function Viewer({booksData}: {booksData: Book[]}) {
  return (
    <>
      <div  className="box-row">
        {booksData.map(({isbn, title, cover}) => (
          <div key={isbn}>
            <div>
              {/* eslint-disable-next-line @next/next/no-img-element */}
              <img src={cover} alt="" width="" height=""></img>
            </div>
          </div>
        ))}
      </div>
      <style jsx>{`
        .box-row {
          display: flex;
          display: -webkit-box;
          display: -ms-flexbox;
          flex-wrap: wrap;
        }
      `}</style>
    </>
  )
}
