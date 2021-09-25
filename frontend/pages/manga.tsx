import { GetStaticProps } from 'next'

type Book = {
	isbn: string
	title: string
	publisher: string
	pubdate: string
	cover: string
	author: string
	subjectCode: string
}

export default function Manga({
  books
}: {
  books: {
    code: number,
    data: Book[]
  }
}) {
  const covers = books.data.map(book => book.cover)
  return (
    <div></div>
  )
}

export const getStaticProps: GetStaticProps = async () => {
  const res = await fetch("http://localhost:8080/backend/manga")
  const books = await res.json()

  return {
    props: {
      books,
    },
  }
}
