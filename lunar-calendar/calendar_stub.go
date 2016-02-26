/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package main

import (
	"pkg.deepin.io/lib/dbus"
)

type Manager struct{}

const (
	LUNAR_DEST = "com.deepin.api.LunarCalendar"
	LUNAR_PATH = "/com/deepin/api/LunarCalendar"
	LUNAR_IFC  = "com.deepin.api.LunarCalendar"
)

func (op *Manager) GetLunarDateBySolar(year, month, day int32) (CaYearInfo, bool, bool) {
	info, ok := getLunarDateBySolar(year, month, day)
	if !ok {
		return CaYearInfo{}, false, false
	}
	leapMonth, _ := getLunarLeapYear(year)
	isLeapMonth := false
	if leapMonth > 0 && leapMonth == info.Month {
		isLeapMonth = true
	} else if leapMonth > 0 && leapMonth > info.Month {
		info.Month += 1
	} else if leapMonth <= 0 {
		info.Month += 1
	}

	logger.Infof("Date: %d - %d - %d\n\tIsLeapMonth: %v",
		info.Year, info.Month, info.Day, isLeapMonth)

	return info, isLeapMonth, true
}

func (op *Manager) GetSolarDateByLunar(year, month, day int32, isLeapMonth bool) (CaYearInfo, bool) {
	leapMonth, _ := getLunarLeapYear(year)
	if leapMonth <= 0 {
		isLeapMonth = false
	}
	if (leapMonth > 0 && month > leapMonth) || isLeapMonth {
		month = month
	} else {
		month -= 1
	}
	if info, ok := lunarToSolar(year, month, day); !ok {
		return CaYearInfo{}, false
	} else {
		return info, true
	}
}

func (op *Manager) GetLunarInfoBySolar(year, month, day int32) (caLunarDayInfo, bool) {
	if info, ok := solarToLunar(year, month, day); !ok {
		return caLunarDayInfo{}, false
	} else {
		return info, true
	}
}

func (op *Manager) GetSolarMonthCalendar(year, month int32, fill bool) (caSolarMonthInfo, bool) {
	logger.Infof("SOLAR DATE: %v- %v- %v", year, month, fill)
	if info, ok := getSolarCalendar(year, month, fill); !ok {
		return caSolarMonthInfo{}, false
	} else {
		logger.Infof("Solar Month Data: %v", info)
		return info, true
	}
}

func (op *Manager) GetLunarMonthCalendar(year, month int32, fill bool) (caLunarMonthInfo, bool) {
	logger.Infof("LUNAR DATE: %v- %v- %v", year, month, fill)
	if info, ok := getLunarCalendar(year, month, fill); !ok {
		return caLunarMonthInfo{}, false
	} else {
		logger.Infof("Lunar Month Data: %v", info)
		return info, true
	}
}

func (op *Manager) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		LUNAR_DEST,
		LUNAR_PATH,
		LUNAR_IFC,
	}
}

func NewManager() *Manager {
	m := &Manager{}

	return m
}
