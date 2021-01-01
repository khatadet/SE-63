import React, { FC } from 'react';
import { HeaderLabel } from '@backstage/core';

const timeFormat = { hour: '2-digit', minute: '2-digit' };
const bkkOptions = { timeZone: 'Asia/Bangkok', ...timeFormat };

const defaultTimes = {
  timeBKK: '',
};

function getTimes() {
  const d = new Date();
  const lang = window.navigator.language;

  // Using the browser native toLocaleTimeString instead of huge moment-tz
  // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toLocaleTimeString
  const timeBKK = d.toLocaleTimeString(lang, bkkOptions);

  return { timeBKK };
}

const HomePageTimer: FC<{}> = () => {
  const [{ timeBKK }, setTimes] = React.useState(
    defaultTimes,
  );

  React.useEffect(() => {
    setTimes(getTimes());

    const intervalId = setInterval(() => {
      setTimes(getTimes());
    }, 1000);

    return () => {
      clearInterval(intervalId);
    };
  }, []);

  return (
    <>
      <HeaderLabel label="BKK" value={timeBKK} />
    </>
  );
};

export default HomePageTimer;
