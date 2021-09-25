import { GetStaticProps } from 'next'
import Viewer from '../components/Viewer'

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
  return (
    <>
    <Viewer booksData={books.data} />
    </>
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
