import {BookData} from '../interfaces'

export default function Viewer({booksData}: {booksData: BookData[]}) {
  return (
    <>
      <div className='box-row'>
        {booksData.map(({isbn, title, cover}) => (
          <div key={isbn}>
            {/* eslint-disable-next-line @next/next/no-img-element */}
            <img src={cover} alt={title} width="" height=""></img>
          </div>
        ))}
      </div>
      <style jsx>{`
        .box-row {
          display: flex;
          display: -webkit-box;
          display: -ms-flexbox;
          flex-wrap: wrap
        }
      `}
      </style>
    </>
  )
}
