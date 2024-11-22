package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/noxscript/ns/v4/effect"
)

func archeryContest(p ns.Player) {
	updateNoxWorldData(p, func(data *NoxWorldData) {
		data.Quest.ArcheryContestScore_Quest01 = 0
	})
	wp := ns.Random(1, 3)
	switch wp {
	case 1:
		ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget01"))
		ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget01"), ns.Waypoint("BarrelTarget01"))
		brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget01"))
		brlhealth := brl.CurrentHealth()
		ns.NewTimer(ns.Frames(44), func() {
			if brlhealth == brl.CurrentHealth() {
				ns.Object("Heckler").ChatStr("That's a MISS!")
			} else {
				ns.Object("Heckler").ChatStr("That's a HIT!")
				updateNoxWorldData(p, func(data *NoxWorldData) {
					data.Quest.ArcheryContestScore_Quest01++
				})
			}
		})
		brl.DeleteAfter(ns.Frames(45))
	case 2:
		ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget02"))
		ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget02"), ns.Waypoint("BarrelTarget02"))
		brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget02"))
		brlhealth := brl.CurrentHealth()
		ns.NewTimer(ns.Frames(44), func() {
			if brlhealth == brl.CurrentHealth() {
				ns.Object("Heckler").ChatStr("That's a MISS!")
			} else {
				ns.Object("Heckler").ChatStr("That's a HIT!")
				updateNoxWorldData(p, func(data *NoxWorldData) {
					data.Quest.ArcheryContestScore_Quest01++
				})
			}
		})
		brl.DeleteAfter(ns.Frames(45))
	case 3:
		ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget03"))
		ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget03"), ns.Waypoint("BarrelTarget03"))
		brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget03"))
		brlhealth := brl.CurrentHealth()
		ns.NewTimer(ns.Frames(44), func() {
			if brlhealth == brl.CurrentHealth() {
				ns.Object("Heckler").ChatStr("That's a MISS!")
			} else {
				ns.Object("Heckler").ChatStr("That's a HIT!")
				updateNoxWorldData(p, func(data *NoxWorldData) {
					data.Quest.ArcheryContestScore_Quest01++
				})
			}
		})
		brl.DeleteAfter(ns.Frames(45))
	}
	ns.NewTimer(ns.Frames(45), func() {
		wp := ns.Random(1, 3)
		switch wp {
		case 1:
			ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget01"))
			ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget01"), ns.Waypoint("BarrelTarget01"))
			brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget01"))
			brlhealth := brl.CurrentHealth()
			ns.NewTimer(ns.Frames(44), func() {
				if brlhealth == brl.CurrentHealth() {
					ns.Object("Heckler").ChatStr("That's a MISS!")
				} else {
					ns.Object("Heckler").ChatStr("That's a HIT!")
					updateNoxWorldData(p, func(data *NoxWorldData) {
						data.Quest.ArcheryContestScore_Quest01++
					})
				}
			})
			brl.DeleteAfter(ns.Frames(45))
		case 2:
			ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget02"))
			ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget02"), ns.Waypoint("BarrelTarget02"))
			brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget02"))
			brlhealth := brl.CurrentHealth()
			ns.NewTimer(ns.Frames(44), func() {
				if brlhealth == brl.CurrentHealth() {
					ns.Object("Heckler").ChatStr("That's a MISS!")
				} else {
					ns.Object("Heckler").ChatStr("That's a HIT!")
					updateNoxWorldData(p, func(data *NoxWorldData) {
						data.Quest.ArcheryContestScore_Quest01++
					})
				}
			})
			brl.DeleteAfter(ns.Frames(45))
		case 3:
			ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget03"))
			ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget03"), ns.Waypoint("BarrelTarget03"))
			brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget03"))
			brlhealth := brl.CurrentHealth()
			ns.NewTimer(ns.Frames(44), func() {
				if brlhealth == brl.CurrentHealth() {
					ns.Object("Heckler").ChatStr("That's a MISS!")
				} else {
					ns.Object("Heckler").ChatStr("That's a HIT!")
					updateNoxWorldData(p, func(data *NoxWorldData) {
						data.Quest.ArcheryContestScore_Quest01++
					})
				}
			})
			brl.DeleteAfter(ns.Frames(45))
		}
		ns.NewTimer(ns.Frames(45), func() {
			wp := ns.Random(1, 3)
			switch wp {
			case 1:
				ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget01"))
				ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget01"), ns.Waypoint("BarrelTarget01"))
				brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget01"))
				brlhealth := brl.CurrentHealth()
				ns.NewTimer(ns.Frames(44), func() {
					if brlhealth == brl.CurrentHealth() {
						ns.Object("Heckler").ChatStr("That's a MISS!")
					} else {
						ns.Object("Heckler").ChatStr("That's a HIT!")
						updateNoxWorldData(p, func(data *NoxWorldData) {
							data.Quest.ArcheryContestScore_Quest01++
						})
					}
				})
				brl.DeleteAfter(ns.Frames(45))
			case 2:
				ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget02"))
				ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget02"), ns.Waypoint("BarrelTarget02"))
				brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget02"))
				brlhealth := brl.CurrentHealth()
				ns.NewTimer(ns.Frames(44), func() {
					if brlhealth == brl.CurrentHealth() {
						ns.Object("Heckler").ChatStr("That's a MISS!")
					} else {
						ns.Object("Heckler").ChatStr("That's a HIT!")
						updateNoxWorldData(p, func(data *NoxWorldData) {
							data.Quest.ArcheryContestScore_Quest01++
						})
					}
				})
				brl.DeleteAfter(ns.Frames(45))
			case 3:
				ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget03"))
				ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget03"), ns.Waypoint("BarrelTarget03"))
				brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget03"))
				brlhealth := brl.CurrentHealth()
				ns.NewTimer(ns.Frames(44), func() {
					if brlhealth == brl.CurrentHealth() {
						ns.Object("Heckler").ChatStr("That's a MISS!")
					} else {
						ns.Object("Heckler").ChatStr("That's a HIT!")
						updateNoxWorldData(p, func(data *NoxWorldData) {
							data.Quest.ArcheryContestScore_Quest01++
						})
					}
				})
				brl.DeleteAfter(ns.Frames(45))
			}
			ns.NewTimer(ns.Frames(45), func() {
				wp := ns.Random(1, 3)
				switch wp {
				case 1:
					ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget01"))
					ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget01"), ns.Waypoint("BarrelTarget01"))
					brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget01"))
					brlhealth := brl.CurrentHealth()
					ns.NewTimer(ns.Frames(44), func() {
						if brlhealth == brl.CurrentHealth() {
							ns.Object("Heckler").ChatStr("That's a MISS!")
						} else {
							ns.Object("Heckler").ChatStr("That's a HIT!")
							updateNoxWorldData(p, func(data *NoxWorldData) {
								data.Quest.ArcheryContestScore_Quest01++
							})
						}
					})
					brl.DeleteAfter(ns.Frames(45))
				case 2:
					ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget02"))
					ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget02"), ns.Waypoint("BarrelTarget02"))
					brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget02"))
					brlhealth := brl.CurrentHealth()
					ns.NewTimer(ns.Frames(44), func() {
						if brlhealth == brl.CurrentHealth() {
							ns.Object("Heckler").ChatStr("That's a MISS!")
						} else {
							ns.Object("Heckler").ChatStr("That's a HIT!")
							updateNoxWorldData(p, func(data *NoxWorldData) {
								data.Quest.ArcheryContestScore_Quest01++
							})
						}
					})
					brl.DeleteAfter(ns.Frames(45))
				case 3:
					ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget03"))
					ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget03"), ns.Waypoint("BarrelTarget03"))
					brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget03"))
					brlhealth := brl.CurrentHealth()
					ns.NewTimer(ns.Frames(44), func() {
						if brlhealth == brl.CurrentHealth() {
							ns.Object("Heckler").ChatStr("That's a MISS!")
						} else {
							ns.Object("Heckler").ChatStr("That's a HIT!")
							updateNoxWorldData(p, func(data *NoxWorldData) {
								data.Quest.ArcheryContestScore_Quest01++
							})
						}
					})
					brl.DeleteAfter(ns.Frames(45))
				}
				ns.NewTimer(ns.Frames(45), func() {
					wp := ns.Random(1, 3)
					switch wp {
					case 1:
						ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget01"))
						ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget01"), ns.Waypoint("BarrelTarget01"))
						brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget01"))
						brlhealth := brl.CurrentHealth()
						ns.NewTimer(ns.Frames(44), func() {
							if brlhealth == brl.CurrentHealth() {
								ns.Object("Heckler").ChatStr("That's a MISS!")
							} else {
								ns.Object("Heckler").ChatStr("That's a HIT!")
								updateNoxWorldData(p, func(data *NoxWorldData) {
									data.Quest.ArcheryContestScore_Quest01++
								})
							}
						})
						brl.DeleteAfter(ns.Frames(45))
					case 2:
						ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget02"))
						ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget02"), ns.Waypoint("BarrelTarget02"))
						brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget02"))
						brlhealth := brl.CurrentHealth()
						ns.NewTimer(ns.Frames(44), func() {
							if brlhealth == brl.CurrentHealth() {
								ns.Object("Heckler").ChatStr("That's a MISS!")
							} else {
								ns.Object("Heckler").ChatStr("That's a HIT!")
								updateNoxWorldData(p, func(data *NoxWorldData) {
									data.Quest.ArcheryContestScore_Quest01++
								})
							}
						})
						brl.DeleteAfter(ns.Frames(45))
					case 3:
						ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget03"))
						ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget03"), ns.Waypoint("BarrelTarget03"))
						brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget03"))
						brlhealth := brl.CurrentHealth()
						ns.NewTimer(ns.Frames(44), func() {
							if brlhealth == brl.CurrentHealth() {
								ns.Object("Heckler").ChatStr("That's a MISS!")
							} else {
								ns.Object("Heckler").ChatStr("That's a HIT!")
								updateNoxWorldData(p, func(data *NoxWorldData) {
									data.Quest.ArcheryContestScore_Quest01++
								})
							}
						})
						brl.DeleteAfter(ns.Frames(45))
					}
					ns.NewTimer(ns.Frames(45), func() {
						wp := ns.Random(1, 3)
						switch wp {
						case 1:
							ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget01"))
							ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget01"), ns.Waypoint("BarrelTarget01"))
							brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget01"))
							brlhealth := brl.CurrentHealth()
							ns.NewTimer(ns.Frames(44), func() {
								if brlhealth == brl.CurrentHealth() {
									ns.Object("Heckler").ChatStr("That's a MISS!")
								} else {
									ns.Object("Heckler").ChatStr("That's a HIT!")
									updateNoxWorldData(p, func(data *NoxWorldData) {
										data.Quest.ArcheryContestScore_Quest01++
									})
								}
							})
							brl.DeleteAfter(ns.Frames(45))
						case 2:
							ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget02"))
							ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget02"), ns.Waypoint("BarrelTarget02"))
							brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget02"))
							brlhealth := brl.CurrentHealth()
							ns.NewTimer(ns.Frames(44), func() {
								if brlhealth == brl.CurrentHealth() {
									ns.Object("Heckler").ChatStr("That's a MISS!")
								} else {
									ns.Object("Heckler").ChatStr("That's a HIT!")
									updateNoxWorldData(p, func(data *NoxWorldData) {
										data.Quest.ArcheryContestScore_Quest01++
									})
								}
							})
							brl.DeleteAfter(ns.Frames(45))
						case 3:
							ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget03"))
							ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget03"), ns.Waypoint("BarrelTarget03"))
							brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget03"))
							brlhealth := brl.CurrentHealth()
							ns.NewTimer(ns.Frames(44), func() {
								if brlhealth == brl.CurrentHealth() {
									ns.Object("Heckler").ChatStr("That's a MISS!")
								} else {
									ns.Object("Heckler").ChatStr("That's a HIT!")
									updateNoxWorldData(p, func(data *NoxWorldData) {
										data.Quest.ArcheryContestScore_Quest01++
									})
								}
							})
							brl.DeleteAfter(ns.Frames(45))
						}
						ns.NewTimer(ns.Frames(45), func() {
							wp := ns.Random(1, 3)
							switch wp {
							case 1:
								ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget01"))
								ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget01"), ns.Waypoint("BarrelTarget01"))
								brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget01"))
								brlhealth := brl.CurrentHealth()
								ns.NewTimer(ns.Frames(44), func() {
									if brlhealth == brl.CurrentHealth() {
										ns.Object("Heckler").ChatStr("That's a MISS!")
									} else {
										ns.Object("Heckler").ChatStr("That's a HIT!")
										updateNoxWorldData(p, func(data *NoxWorldData) {
											data.Quest.ArcheryContestScore_Quest01++
										})
									}
								})
								brl.DeleteAfter(ns.Frames(45))
							case 2:
								ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget02"))
								ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget02"), ns.Waypoint("BarrelTarget02"))
								brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget02"))
								brlhealth := brl.CurrentHealth()
								ns.NewTimer(ns.Frames(44), func() {
									if brlhealth == brl.CurrentHealth() {
										ns.Object("Heckler").ChatStr("That's a MISS!")
									} else {
										ns.Object("Heckler").ChatStr("That's a HIT!")
										updateNoxWorldData(p, func(data *NoxWorldData) {
											data.Quest.ArcheryContestScore_Quest01++
										})
									}
								})
								brl.DeleteAfter(ns.Frames(45))
							case 3:
								ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget03"))
								ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget03"), ns.Waypoint("BarrelTarget03"))
								brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget03"))
								brlhealth := brl.CurrentHealth()
								ns.NewTimer(ns.Frames(44), func() {
									if brlhealth == brl.CurrentHealth() {
										ns.Object("Heckler").ChatStr("That's a MISS!")
									} else {
										ns.Object("Heckler").ChatStr("That's a HIT!")
										updateNoxWorldData(p, func(data *NoxWorldData) {
											data.Quest.ArcheryContestScore_Quest01++
										})
									}
								})
								brl.DeleteAfter(ns.Frames(45))
							}
							ns.NewTimer(ns.Frames(45), func() {
								wp := ns.Random(1, 3)
								switch wp {
								case 1:
									ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget01"))
									ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget01"), ns.Waypoint("BarrelTarget01"))
									brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget01"))
									brlhealth := brl.CurrentHealth()
									ns.NewTimer(ns.Frames(44), func() {
										if brlhealth == brl.CurrentHealth() {
											ns.Object("Heckler").ChatStr("That's a MISS!")
										} else {
											ns.Object("Heckler").ChatStr("That's a HIT!")
											updateNoxWorldData(p, func(data *NoxWorldData) {
												data.Quest.ArcheryContestScore_Quest01++
											})
										}
									})
									brl.DeleteAfter(ns.Frames(45))
								case 2:
									ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget02"))
									ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget02"), ns.Waypoint("BarrelTarget02"))
									brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget02"))
									brlhealth := brl.CurrentHealth()
									ns.NewTimer(ns.Frames(44), func() {
										if brlhealth == brl.CurrentHealth() {
											ns.Object("Heckler").ChatStr("That's a MISS!")
										} else {
											ns.Object("Heckler").ChatStr("That's a HIT!")
											updateNoxWorldData(p, func(data *NoxWorldData) {
												data.Quest.ArcheryContestScore_Quest01++
											})
										}
									})
									brl.DeleteAfter(ns.Frames(45))
								case 3:
									ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget03"))
									ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget03"), ns.Waypoint("BarrelTarget03"))
									brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget03"))
									brlhealth := brl.CurrentHealth()
									ns.NewTimer(ns.Frames(44), func() {
										if brlhealth == brl.CurrentHealth() {
											ns.Object("Heckler").ChatStr("That's a MISS!")
										} else {
											ns.Object("Heckler").ChatStr("That's a HIT!")
											updateNoxWorldData(p, func(data *NoxWorldData) {
												data.Quest.ArcheryContestScore_Quest01++
											})
										}
									})
									brl.DeleteAfter(ns.Frames(45))
								}
								ns.NewTimer(ns.Frames(45), func() {
									wp := ns.Random(1, 3)
									switch wp {
									case 1:
										ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget01"))
										ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget01"), ns.Waypoint("BarrelTarget01"))
										brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget01"))
										brlhealth := brl.CurrentHealth()
										ns.NewTimer(ns.Frames(44), func() {
											if brlhealth == brl.CurrentHealth() {
												ns.Object("Heckler").ChatStr("That's a MISS!")
											} else {
												ns.Object("Heckler").ChatStr("That's a HIT!")
												updateNoxWorldData(p, func(data *NoxWorldData) {
													data.Quest.ArcheryContestScore_Quest01++
												})
											}
										})
										brl.DeleteAfter(ns.Frames(45))
									case 2:
										ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget02"))
										ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget02"), ns.Waypoint("BarrelTarget02"))
										brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget02"))
										brlhealth := brl.CurrentHealth()
										ns.NewTimer(ns.Frames(44), func() {
											if brlhealth == brl.CurrentHealth() {
												ns.Object("Heckler").ChatStr("That's a MISS!")
											} else {
												ns.Object("Heckler").ChatStr("That's a HIT!")
												updateNoxWorldData(p, func(data *NoxWorldData) {
													data.Quest.ArcheryContestScore_Quest01++
												})
											}
										})
										brl.DeleteAfter(ns.Frames(45))
									case 3:
										ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget03"))
										ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget03"), ns.Waypoint("BarrelTarget03"))
										brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget03"))
										brlhealth := brl.CurrentHealth()
										ns.NewTimer(ns.Frames(44), func() {
											if brlhealth == brl.CurrentHealth() {
												ns.Object("Heckler").ChatStr("That's a MISS!")
											} else {
												ns.Object("Heckler").ChatStr("That's a HIT!")
												updateNoxWorldData(p, func(data *NoxWorldData) {
													data.Quest.ArcheryContestScore_Quest01++
												})
											}
										})
										brl.DeleteAfter(ns.Frames(45))
									}
									ns.NewTimer(ns.Frames(45), func() {
										wp := ns.Random(1, 3)
										switch wp {
										case 1:
											ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget01"))
											ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget01"), ns.Waypoint("BarrelTarget01"))
											brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget01"))
											brlhealth := brl.CurrentHealth()
											ns.NewTimer(ns.Frames(44), func() {
												if brlhealth == brl.CurrentHealth() {
													ns.Object("Heckler").ChatStr("That's a MISS!")
												} else {
													ns.Object("Heckler").ChatStr("That's a HIT!")
													updateNoxWorldData(p, func(data *NoxWorldData) {
														data.Quest.ArcheryContestScore_Quest01++
													})
												}
											})
											brl.DeleteAfter(ns.Frames(45))
										case 2:
											ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget02"))
											ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget02"), ns.Waypoint("BarrelTarget02"))
											brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget02"))
											brlhealth := brl.CurrentHealth()
											ns.NewTimer(ns.Frames(44), func() {
												if brlhealth == brl.CurrentHealth() {
													ns.Object("Heckler").ChatStr("That's a MISS!")
												} else {
													ns.Object("Heckler").ChatStr("That's a HIT!")
													updateNoxWorldData(p, func(data *NoxWorldData) {
														data.Quest.ArcheryContestScore_Quest01++
													})
												}
											})
											brl.DeleteAfter(ns.Frames(45))
										case 3:
											ns.AudioEvent(audio.BlinkCast, ns.Waypoint("BarrelTarget03"))
											ns.Effect(effect.SMOKE_BLAST, ns.Waypoint("BarrelTarget03"), ns.Waypoint("BarrelTarget03"))
											brl := ns.CreateObject("TargetBarrel1", ns.Waypoint("BarrelTarget03"))
											brlhealth := brl.CurrentHealth()
											ns.NewTimer(ns.Frames(44), func() {
												if brlhealth == brl.CurrentHealth() {
													ns.Object("Heckler").ChatStr("That's a MISS!")
												} else {
													ns.Object("Heckler").ChatStr("That's a HIT!")
													updateNoxWorldData(p, func(data *NoxWorldData) {
														data.Quest.ArcheryContestScore_Quest01++
													})
												}
											})
											brl.DeleteAfter(ns.Frames(45))
										}
										ns.NewTimer(ns.Frames(45), func() {
											ArcheryContestActive = false
											data := loadMyNoxWorldData(p)
											switch data.Quest.ArcheryContestScore_Quest01 {
											case 0:
												ns.Object("Heckler").ChatStr("You didn't hit any! It's a good day to be a barrel, I guess.")
											case 1:
												ns.Object("Heckler").ChatStr("You hit 1 out of 10. Try not to hurt yourself with that bow.")
											case 2:
												ns.Object("Heckler").ChatStr("You hit 2 out of 10. I guess luck counts for something.")
											case 3:
												ns.Object("Heckler").ChatStr("You hit 3 out of 10. Maybe we should get bigger barrels for you.")
											case 4:
												ns.Object("Heckler").ChatStr("You hit 4 out of 10. Maybe you need some spectacles.")
											case 5:
												ns.Object("Heckler").ChatStr("You hit 5 out of 10. Well, you hit half of them -- but you missed half too.")
											case 6:
												ns.Object("Heckler").ChatStr("You hit 6 out of 10. Maybe a drink would steady your hand.")
											case 7:
												ns.Object("Heckler").ChatStr("You hit 7 out of 10. A fair effort.")
											case 8:
												ns.Object("Heckler").ChatStr("You hit 8 out of 10. Good shooting, but not quite good enough.")
											case 9:
												ns.Object("Heckler").ChatStr("Excellent shooting... You hit 9 out of 10!")
											case 10:
												ns.Object("Heckler").ChatStr("Exceptional shooting... You hit 10 out of 10!")
											}
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})
}
